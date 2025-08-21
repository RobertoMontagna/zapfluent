package config_test

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
)

const (
	testFieldName    = "test-field"
	testErrorMessage = "test-error"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := config.NewErrorHandlingConfiguration()

		g.Expect(cfg.Mode()).To(Equal(config.ErrorHandlingModeContinue))
	})

	t.Run("with WithMode option", func(t *testing.T) {
		opt := config.WithMode(config.ErrorHandlingModeEarlyFailing)

		cfg := config.NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.Mode()).To(Equal(config.ErrorHandlingModeEarlyFailing))
	})

	t.Run("with WithFallbackErrorMessage option", func(t *testing.T) {
		const message = "test-message"
		opt := config.WithFallbackErrorMessage(message)

		cfg := config.NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.FallbackErrorMessage).To(Equal(message))
	})
}

func TestErrorHandlingMode_String(t *testing.T) {
	g := NewWithT(t)

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

			g.Expect(s).To(Equal(tc.expected))
		})
	}
}

func TestIntToErrorHandlingMode(t *testing.T) {
	g := NewWithT(t)

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

			g.Expect(mode).To(Equal(tc.expected))
		})
	}
}

func TestFixedStringFallback(t *testing.T) {
	g := NewWithT(t)
	const fallbackValue = "fixed-value"
	factory := config.FixedStringFallback(fallbackValue)

	field := factory(testFieldName, errors.New(testErrorMessage))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, fallbackValue))
}

func TestErrorStringFallback(t *testing.T) {
	g := NewWithT(t)
	const errorMsg = "this is the error message"
	factory := config.ErrorStringFallback()

	field := factory(testFieldName, errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, errorMsg))
}
