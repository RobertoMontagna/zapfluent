package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/testutil"
)

const (
	testError1        = "error 1"
	testError2        = "error 2"
	testFieldName1    = "field1"
	testFieldName2    = "field2"
	testOriginalError = "original error"
	testFallbackValue = "fallback-value"
	testFailingField  = "failing_field"
	testFallbackError = "fallback failed"
)

func newFluentWithConfig(cfg config.Configuration) (*zapfluent.Fluent, *zapcore.MapObjectEncoder) {
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.NewFluent(enc, cfg)
	return fluent, enc
}

func TestFluent_Done_WithMultipleErrors_AggregatesErrors(t *testing.T) {
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)
	field1 := testutil.FailingField{Err: err1}
	field2 := testutil.FailingField{Err: err2}

	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.NewFluent(enc, config.NewConfiguration())

	err := fluent.Add(field1).Add(field2).Done()

	assert.ErrorContains(t, err, testError1)
	assert.ErrorContains(t, err, testError2)
}

func TestFluent_errorHandling_EarlyFailing(t *testing.T) {
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)
	field1 := testutil.FailingField{Err: err1, NameValue: testFieldName1}
	field2 := testutil.FailingField{Err: err2, NameValue: testFieldName2}

	cfg := config.NewConfiguration(
		config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		),
	)
	fluent, _ := newFluentWithConfig(cfg)

	err := fluent.Add(field1).Add(field2).Done()

	assert.Equal(t, err1, err) // Should only return the first error
}

func TestFluent_WithFallback(t *testing.T) {
	t.Run("it replaces the failing field and aggregates the error", func(t *testing.T) {
		testErr := errors.New(testOriginalError)
		cfg := config.NewConfiguration(
			config.WithErrorHandling(
				config.NewErrorHandlingConfiguration(
					config.WithFallbackFieldFactory(config.FixedStringFallback(testFallbackValue)),
				),
			),
		)
		fluent, enc := newFluentWithConfig(cfg)
		failingField := testutil.FailingField{Err: testErr, NameValue: testFailingField}

		err := fluent.Add(failingField).Done()

		assert.Equal(t, testErr, err, "The original error should be aggregated")

		fallbackValue, exists := enc.Fields[testFailingField]
		assert.True(t, exists, "The fallback field should have been added")
		assert.Equal(t, testFallbackValue, fallbackValue)
	})

	t.Run("it aggregates errors from the fallback field itself", func(t *testing.T) {
		originalErr := errors.New(testOriginalError)
		fallbackErr := errors.New(testFallbackError)

		failingFactory := func(name string, err error) fluentfield.Field {
			return testutil.FailingField{NameValue: name, Err: fallbackErr}
		}

		cfg := config.NewConfiguration(
			config.WithErrorHandling(
				config.NewErrorHandlingConfiguration(
					config.WithFallbackFieldFactory(failingFactory),
				),
			),
		)
		fluent, enc := newFluentWithConfig(cfg)
		initialFailingField := testutil.FailingField{Err: originalErr, NameValue: testFailingField}

		err := fluent.Add(initialFailingField).Done()

		assert.ErrorIs(t, err, originalErr, "The original error should be aggregated")
		assert.ErrorIs(t, err, fallbackErr, "The fallback's error should also be aggregated")
		assert.Empty(t, enc.Fields, "No field should have been successfully encoded")
	})
}
