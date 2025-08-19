package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

type errorField struct {
	err  error
	name string
}

func (f errorField) Encode(enc zapcore.ObjectEncoder) error {
	return f.err
}

func (f errorField) Name() string {
	if f.name == "" {
		return "error"
	}
	return f.name
}

func TestFluent_errorHandling(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := errorField{err: err1}
	field2 := errorField{err: err2}
	fluent := zapfluent.AsFluent(nil)

	err := fluent.Add(field1).Add(field2).Done()

	assert.ErrorContains(t, err, "error 1")
	assert.ErrorContains(t, err, "error 2")
}

func TestFluent_errorHandling_EarlyFailing(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := errorField{err: err1, name: "field1"}
	field2 := errorField{err: err2, name: "field2"}

	cfg := config.NewConfiguration(
		config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		),
	)

	// Create a dummy encoder
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	// Create a fluent instance with the custom config
	fluent := zapfluent.NewFluent(enc, cfg)

	err := fluent.Add(field1).Add(field2).Done()

	assert.Equal(t, err1, err) // Should only return the first error
}

func TestFluent_WithFallback(t *testing.T) {
	t.Run("it replaces the failing field and aggregates the error", func(t *testing.T) {
		// Arrange
		enc := zapcore.NewMapObjectEncoder()
		testErr := errors.New("original error")
		cfg := config.NewConfiguration(
			config.WithErrorHandling(
				config.NewErrorHandlingConfiguration(
					config.WithFallbackFieldFactory(config.FixedStringFallback("fallback-value")),
				),
			),
		)
		fluent := zapfluent.NewFluent(enc, cfg)
		failingField := errorField{err: testErr, name: "failing_field"}

		// Act
		err := fluent.Add(failingField).Done()

		// Assert
		assert.Equal(t, testErr, err, "The original error should be aggregated")

		fallbackValue, exists := enc.Fields["failing_field"]
		assert.True(t, exists, "The fallback field should have been added")
		assert.Equal(t, "fallback-value", fallbackValue)
	})

	t.Run("it aggregates errors from the fallback field itself", func(t *testing.T) {
		// Arrange
		enc := zapcore.NewMapObjectEncoder()
		originalErr := errors.New("original error")
		fallbackErr := errors.New("fallback failed")

		failingFactory := func(name string, err error) fluentfield.Field {
			return errorField{name: name, err: fallbackErr}
		}

		cfg := config.NewConfiguration(
			config.WithErrorHandling(
				config.NewErrorHandlingConfiguration(
					config.WithFallbackFieldFactory(failingFactory),
				),
			),
		)
		fluent := zapfluent.NewFluent(enc, cfg)
		initialFailingField := errorField{err: originalErr, name: "failing_field"}

		// Act
		err := fluent.Add(initialFailingField).Done()

		// Assert
		assert.ErrorIs(t, err, originalErr, "The original error should be aggregated")
		assert.ErrorIs(t, err, fallbackErr, "The fallback's error should also be aggregated")
		assert.Empty(t, enc.Fields, "No field should have been successfully encoded")
	})
}
