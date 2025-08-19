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

func newFluentWithConfig(cfg config.Configuration) (*zapfluent.Fluent, *zapcore.MapObjectEncoder) {
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.NewFluent(enc, cfg)
	return fluent, enc
}

func TestFluent_Done_WithMultipleErrors_AggregatesErrors(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := testutil.FailingField{Err: err1}
	field2 := testutil.FailingField{Err: err2}

	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.NewFluent(enc, config.NewConfiguration())

	err := fluent.Add(field1).Add(field2).Done()

	assert.ErrorContains(t, err, "error 1")
	assert.ErrorContains(t, err, "error 2")
}

func TestFluent_errorHandling_EarlyFailing(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := testutil.FailingField{Err: err1, NameValue: "field1"}
	field2 := testutil.FailingField{Err: err2, NameValue: "field2"}

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
		testErr := errors.New("original error")
		cfg := config.NewConfiguration(
			config.WithErrorHandling(
				config.NewErrorHandlingConfiguration(
					config.WithFallbackFieldFactory(config.FixedStringFallback("fallback-value")),
				),
			),
		)
		fluent, enc := newFluentWithConfig(cfg)
		failingField := testutil.FailingField{Err: testErr, NameValue: "failing_field"}

		err := fluent.Add(failingField).Done()

		assert.Equal(t, testErr, err, "The original error should be aggregated")

		fallbackValue, exists := enc.Fields["failing_field"]
		assert.True(t, exists, "The fallback field should have been added")
		assert.Equal(t, "fallback-value", fallbackValue)
	})

	t.Run("it aggregates errors from the fallback field itself", func(t *testing.T) {
		originalErr := errors.New("original error")
		fallbackErr := errors.New("fallback failed")

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
		initialFailingField := testutil.FailingField{Err: originalErr, NameValue: "failing_field"}

		err := fluent.Add(initialFailingField).Done()

		assert.ErrorIs(t, err, originalErr, "The original error should be aggregated")
		assert.ErrorIs(t, err, fallbackErr, "The fallback's error should also be aggregated")
		assert.Empty(t, enc.Fields, "No field should have been successfully encoded")
	})
}
