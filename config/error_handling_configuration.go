package config

import (
	"go.robertomontagna.dev/zapfluent/enum"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

// FallbackFieldFactory is a function that creates a fallback field.
// It receives the name of the field that failed to encode and the error that occurred.
type FallbackFieldFactory func(name string, err error) fluentfield.Field

// An ErrorHandlingConfigurationOption is a function that applies a configuration
// to an ErrorHandlingConfiguration object.
type ErrorHandlingConfigurationOption func(*ErrorHandlingConfiguration)

// WithMode is an ErrorHandlingConfigurationOption that sets the error handling
// mode (e.g., continue on error or fail early).
func WithMode(mode ErrorHandlingMode) ErrorHandlingConfigurationOption {
	return func(c *ErrorHandlingConfiguration) {
		c.mode = mode
	}
}

// WithFallbackFieldFactory is an ErrorHandlingConfigurationOption that sets a
// factory function to produce a fallback field when an encoding error occurs.
func WithFallbackFieldFactory(factory FallbackFieldFactory) ErrorHandlingConfigurationOption {
	return func(c *ErrorHandlingConfiguration) {
		c.fallbackFactory = optional.Some(factory)
	}
}

// NewErrorHandlingConfiguration creates a new ErrorHandlingConfiguration with
// the given options.
//
// If no options are provided, it returns a default configuration that continues
// on error and does not use a fallback factory.
func NewErrorHandlingConfiguration(opts ...ErrorHandlingConfigurationOption) ErrorHandlingConfiguration {
	config := ErrorHandlingConfiguration{
		mode:            ErrorHandlingModeContinue,
		fallbackFactory: optional.Empty[FallbackFieldFactory](),
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

// ErrorHandlingConfiguration holds the settings that define how errors are
// handled during field encoding.
type ErrorHandlingConfiguration struct {
	mode            ErrorHandlingMode
	fallbackFactory optional.Optional[FallbackFieldFactory]
}

// Mode returns the configured error handling mode.
func (c ErrorHandlingConfiguration) Mode() ErrorHandlingMode {
	return c.mode
}

// FallbackFactory returns an optional factory function for creating fallback
// fields. The optional will be empty if no factory is configured.
func (c ErrorHandlingConfiguration) FallbackFactory() optional.Optional[FallbackFieldFactory] {
	return c.fallbackFactory
}

// ErrorHandlingMode defines the strategy for handling errors that occur during
// field encoding.
type ErrorHandlingMode int

const (
	// ErrorHandlingModeUnknown is the default, zero-value handling mode.
	ErrorHandlingModeUnknown ErrorHandlingMode = iota
	// ErrorHandlingModeEarlyFailing stops processing fields as soon as the
	// first error is encountered. The final aggregated error will contain
	// only the first error that occurred.
	ErrorHandlingModeEarlyFailing
	// ErrorHandlingModeContinue continues processing fields even after errors
	// occur. All errors are collected and returned as a single aggregated
	// error. This is the default behavior.
	ErrorHandlingModeContinue
)

const (
	ErrorHandlingModeUnknownString      = "Unknown"
	ErrorHandlingModeEarlyFailingString = "EarlyFailing"
	ErrorHandlingModeContinueString     = "Continue"
)

var errorHandlingModeEnum = enum.New(
	map[ErrorHandlingMode]string{
		ErrorHandlingModeUnknown:      ErrorHandlingModeUnknownString,
		ErrorHandlingModeEarlyFailing: ErrorHandlingModeEarlyFailingString,
		ErrorHandlingModeContinue:     ErrorHandlingModeContinueString,
	},
	ErrorHandlingModeUnknown,
)

// String returns the string representation of the ErrorHandlingMode.
func (m ErrorHandlingMode) String() string {
	return errorHandlingModeEnum.String(m)
}

// IntToErrorHandlingMode converts an integer to an ErrorHandlingMode.
// If the integer does not correspond to a valid mode, it returns
// ErrorHandlingModeUnknown.
func IntToErrorHandlingMode(value int) ErrorHandlingMode {
	return errorHandlingModeEnum.FromInt(value)
}

// FixedStringFallback returns a FallbackFieldFactory that creates a field with a
// predefined, fixed string value.
func FixedStringFallback(value string) FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, value)
	}
}

// ErrorStringFallback returns a FallbackFieldFactory that creates a field whose
// value is the string representation of the error that occurred.
func ErrorStringFallback() FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, err.Error())
	}
}
