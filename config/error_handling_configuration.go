package config

import (
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

// FallbackFieldFactory is a function that creates a fallback field.
// It receives the name of the field that failed to encode and the error that occurred.
type FallbackFieldFactory func(name string, err error) fluentfield.Field

type ErrorHandlingConfigurationOption func(*ErrorHandlingConfiguration)

func WithMode(mode ErrorHandlingMode) ErrorHandlingConfigurationOption {
	return func(c *ErrorHandlingConfiguration) {
		c.mode = mode
	}
}

func WithFallbackFieldFactory(factory FallbackFieldFactory) ErrorHandlingConfigurationOption {
	return func(c *ErrorHandlingConfiguration) {
		c.fallbackFactory = optional.Some(factory)
	}
}

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

type ErrorHandlingConfiguration struct {
	mode            ErrorHandlingMode
	fallbackFactory optional.Optional[FallbackFieldFactory]
}

func (c ErrorHandlingConfiguration) Mode() ErrorHandlingMode {
	return c.mode
}

func (c ErrorHandlingConfiguration) FallbackFactory() optional.Optional[FallbackFieldFactory] {
	return c.fallbackFactory
}
