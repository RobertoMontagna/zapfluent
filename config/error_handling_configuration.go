package config

type ErrorHandlingConfigurationOption func(*ErrorHandlingConfiguration)

func WithMode(mode ErrorHandlingMode) ErrorHandlingConfigurationOption {
	return func(c *ErrorHandlingConfiguration) {
		c.mode = mode
	}
}

func NewErrorHandlingConfiguration(opts ...ErrorHandlingConfigurationOption) ErrorHandlingConfiguration {
	config := ErrorHandlingConfiguration{
		mode: ErrorHandlingModeContinue,
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

type ErrorHandlingConfiguration struct {
	mode ErrorHandlingMode
}

func (c ErrorHandlingConfiguration) Mode() ErrorHandlingMode {
	return c.mode
}
