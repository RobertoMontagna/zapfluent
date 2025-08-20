package config_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
)

const (
	testFieldName    = "test-field"
	testErrorMessage = "test-error"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	t.Run("with default options", func(t *testing.T) {
		cfg := config.NewErrorHandlingConfiguration()

		assert.Equal(t, config.ErrorHandlingModeContinue, cfg.Mode())
	})

	t.Run("with WithMode option", func(t *testing.T) {
		opt := config.WithMode(config.ErrorHandlingModeEarlyFailing)

		cfg := config.NewErrorHandlingConfiguration(opt)

		assert.Equal(t, config.ErrorHandlingModeEarlyFailing, cfg.Mode())
	})
}

func TestErrorHandlingMode_String(t *testing.T) {
	testCases := []struct {
		mode     config.ErrorHandlingMode
		expected string
	}{
		{config.ErrorHandlingModeUnknown, config.ErrorHandlingModeUnknownString},
		{config.ErrorHandlingModeEarlyFailing, config.ErrorHandlingModeEarlyFailingString},
		{config.ErrorHandlingModeContinue, config.ErrorHandlingModeContinueString},
		{config.ErrorHandlingMode(99), "Unknown(99)"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			s := tc.mode.String()

			assert.Equal(t, tc.expected, s)
		})
	}
}

func TestIntToErrorHandlingMode(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected config.ErrorHandlingMode
	}{
		{config.ErrorHandlingModeUnknownString, 0, config.ErrorHandlingModeUnknown},
		{config.ErrorHandlingModeEarlyFailingString, 1, config.ErrorHandlingModeEarlyFailing},
		{config.ErrorHandlingModeContinueString, 2, config.ErrorHandlingModeContinue},
		{"Invalid", 99, config.ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mode := config.IntToErrorHandlingMode(tc.value)

			assert.Equal(t, tc.expected, mode)
		})
	}
}

func TestFixedStringFallback(t *testing.T) {
	const fallbackValue = "fixed-value"
	factory := config.FixedStringFallback(fallbackValue)

	field := factory(testFieldName, errors.New(testErrorMessage))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, fallbackValue, enc.Fields[testFieldName])
}

func TestErrorStringFallback(t *testing.T) {
	const errorMsg = "this is the error message"
	factory := config.ErrorStringFallback()

	field := factory(testFieldName, errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, errorMsg, enc.Fields[testFieldName])
}
