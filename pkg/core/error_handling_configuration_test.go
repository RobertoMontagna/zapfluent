package core_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

const (
	ehcTestFieldName    = "test-field"
	ehcTestErrorMessage = "test-error"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	const customMessage = "test-message"

	testCases := []struct {
		name                   string
		options                []core.ErrorHandlingConfigurationOption
		expectedMode           core.ErrorHandlingMode
		expectedFallbackErrMsg string
	}{
		{
			name:                   "with default options",
			options:                []core.ErrorHandlingConfigurationOption{},
			expectedMode:           core.ErrorHandlingModeContinue,
			expectedFallbackErrMsg: "failed to encode fallback field",
		},
		{
			name: "with WithMode option",
			options: []core.ErrorHandlingConfigurationOption{
				core.WithMode(core.ErrorHandlingModeEarlyFailing),
			},
			expectedMode:           core.ErrorHandlingModeEarlyFailing,
			expectedFallbackErrMsg: "failed to encode fallback field",
		},
		{
			name: "with WithFallbackErrorMessage option",
			options: []core.ErrorHandlingConfigurationOption{
				core.WithFallbackErrorMessage(customMessage),
			},
			expectedMode:           core.ErrorHandlingModeContinue,
			expectedFallbackErrMsg: customMessage,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			cfg := core.NewErrorHandlingConfiguration(tc.options...)

			g.Expect(cfg.Mode()).To(Equal(tc.expectedMode))
			g.Expect(cfg.FallbackErrorMessage).To(Equal(tc.expectedFallbackErrMsg))
		})
	}
}

func TestErrorHandlingMode_String(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		mode     core.ErrorHandlingMode
		expected string
	}{
		{core.ErrorHandlingModeUnknown, core.ErrorHandlingModeUnknownString},
		{core.ErrorHandlingModeEarlyFailing, core.ErrorHandlingModeEarlyFailingString},
		{core.ErrorHandlingModeContinue, core.ErrorHandlingModeContinueString},
		{core.ErrorHandlingMode(99), "Unknown(99)"},
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
		expected core.ErrorHandlingMode
	}{
		{core.ErrorHandlingModeUnknownString, 0, core.ErrorHandlingModeUnknown},
		{core.ErrorHandlingModeEarlyFailingString, 1, core.ErrorHandlingModeEarlyFailing},
		{core.ErrorHandlingModeContinueString, 2, core.ErrorHandlingModeContinue},
		{"Invalid", 99, core.ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mode := core.IntToErrorHandlingMode(tc.value)

			g.Expect(mode).To(Equal(tc.expected))
		})
	}
}

func TestFixedStringFallback_WhenCalled_ReturnsFieldWithFixedValue(t *testing.T) {
	g := NewWithT(t)
	const fallbackValue = "fixed-value"

	factory := core.FixedStringFallback(fallbackValue)
	enc := zapcore.NewMapObjectEncoder()

	field := factory(ehcTestFieldName, errors.New(ehcTestErrorMessage))
	err := field.Encode(enc)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(ehcTestFieldName, fallbackValue))
}

func TestErrorStringFallback_WhenCalled_ReturnsFieldWithErrorString(t *testing.T) {
	g := NewWithT(t)
	const errorMsg = "this is the error message"

	factory := core.ErrorStringFallback()
	enc := zapcore.NewMapObjectEncoder()

	field := factory(ehcTestFieldName, errors.New(errorMsg))
	err := field.Encode(enc)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(ehcTestFieldName, errorMsg))
}
