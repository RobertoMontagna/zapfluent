package core

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	. "github.com/onsi/gomega"
)

const (
	testFieldName    = "test-field"
	testErrorMessage = "test-error"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := NewErrorHandlingConfiguration()

		g.Expect(cfg.Mode()).To(Equal(ErrorHandlingModeContinue))
	})

	t.Run("with WithMode option", func(t *testing.T) {
		opt := WithMode(ErrorHandlingModeEarlyFailing)

		cfg := NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.Mode()).To(Equal(ErrorHandlingModeEarlyFailing))
	})

	t.Run("with WithFallbackErrorMessage option", func(t *testing.T) {
		const message = "test-message"
		opt := WithFallbackErrorMessage(message)

		cfg := NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.FallbackErrorMessage).To(Equal(message))
	})
}

func TestErrorHandlingMode_String(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		mode     ErrorHandlingMode
		expected string
	}{
		{ErrorHandlingModeUnknown, ErrorHandlingModeUnknownString},
		{ErrorHandlingModeEarlyFailing, ErrorHandlingModeEarlyFailingString},
		{ErrorHandlingModeContinue, ErrorHandlingModeContinueString},
		{ErrorHandlingMode(99), "Unknown(99)"},
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
		expected ErrorHandlingMode
	}{
		{ErrorHandlingModeUnknownString, 0, ErrorHandlingModeUnknown},
		{ErrorHandlingModeEarlyFailingString, 1, ErrorHandlingModeEarlyFailing},
		{ErrorHandlingModeContinueString, 2, ErrorHandlingModeContinue},
		{"Invalid", 99, ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mode := IntToErrorHandlingMode(tc.value)

			g.Expect(mode).To(Equal(tc.expected))
		})
	}
}

func TestFixedStringFallback(t *testing.T) {
	g := NewWithT(t)
	const fallbackValue = "fixed-value"
	factory := FixedStringFallback(fallbackValue)

	field := factory(testFieldName, errors.New(testErrorMessage))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, fallbackValue))
}

func TestErrorStringFallback(t *testing.T) {
	g := NewWithT(t)
	const errorMsg = "this is the error message"
	factory := ErrorStringFallback()

	field := factory(testFieldName, errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, errorMsg))
}
