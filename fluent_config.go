package zapfluent

type FluentConfigOption func(*FluentConfig)

func WithErrorManagementMode(mode ErrorManagementMode) FluentConfigOption {
	return func(c *FluentConfig) {
		c.errorManagementMode = mode
	}
}

// NewFluentConfig creates a new FluentConfig with the given options.
// By default, the error management mode is ErrorManagementModeContinuePrintFailures.
func NewFluentConfig(opts ...FluentConfigOption) FluentConfig {
	config := FluentConfig{
		errorManagementMode: ErrorManagementModeContinuePrintFailures,
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

type FluentConfig struct {
	errorManagementMode ErrorManagementMode
}

func (c FluentConfig) ErrorManagementMode() ErrorManagementMode {
	return c.errorManagementMode
}

func (c FluentConfig) Clone(opts ...FluentConfigOption) FluentConfig {
	newConfig := c
	for _, opt := range opts {
		opt(&newConfig)
	}
	return newConfig
}
