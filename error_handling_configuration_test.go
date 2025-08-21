package zapfluent_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"

	. "github.com/onsi/gomega"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := zapfluent.NewErrorHandlingConfiguration()

		g.Expect(cfg.Mode()).To(Equal(zapfluent.ErrorHandlingModeContinue))
	})

	t.Run("with WithMode option", func(t *testing.T) {
		opt := zapfluent.WithMode(zapfluent.ErrorHandlingModeEarlyFailing)

		cfg := zapfluent.NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.Mode()).To(Equal(zapfluent.ErrorHandlingModeEarlyFailing))
	})

	t.Run("with WithFallbackErrorMessage option", func(t *testing.T) {
		const message = "test-message"
		opt := zapfluent.WithFallbackErrorMessage(message)

		cfg := zapfluent.NewErrorHandlingConfiguration(opt)

		g.Expect(cfg.FallbackErrorMessage).To(Equal(message))
	})
}

func TestErrorHandlingMode_String(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		mode     zapfluent.ErrorHandlingMode
		expected string
	}{
		{zapfluent.ErrorHandlingModeUnknown, "Unknown"},
		{zapfluent.ErrorHandlingModeEarlyFailing, "EarlyFailing"},
		{zapfluent.ErrorHandlingModeContinue, "Continue"},
		{zapfluent.ErrorHandlingMode(99), "Unknown(99)"},
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
		expected zapfluent.ErrorHandlingMode
	}{
		{"Unknown", 0, zapfluent.ErrorHandlingModeUnknown},
		{"EarlyFailing", 1, zapfluent.ErrorHandlingModeEarlyFailing},
		{"Continue", 2, zapfluent.ErrorHandlingModeContinue},
		{"Invalid", 99, zapfluent.ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mode := zapfluent.IntToErrorHandlingMode(tc.value)

			g.Expect(mode).To(Equal(tc.expected))
		})
	}
}

func TestFixedStringFallback(t *testing.T) {
	g := NewWithT(t)
	const fallbackValue = "fixed-value"
	factory := zapfluent.FixedStringFallback(fallbackValue)

	field := factory("test-field", errors.New("test-error"))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(enc.Fields).To(HaveKeyWithValue("test-field", fallbackValue))
}

func TestErrorStringFallback(t *testing.T) {
	g := NewWithT(t)
	const errorMsg = "this is the error message"
	factory := zapfluent.ErrorStringFallback()

	field := factory("test-field", errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(enc.Fields).To(HaveKeyWithValue("test-field", errorMsg))
}
