package config

type ConfigurationOption func(*Configuration)

func WithErrorHandling(opts ...ErrorHandlingConfigurationOption) ConfigurationOption {
	return func(c *Configuration) {
		c.errorHandling = NewErrorHandlingConfiguration(opts...)
	}
}

func NewConfiguration(opts ...ConfigurationOption) Configuration {
	config := Configuration{
		errorHandling: NewErrorHandlingConfiguration(),
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

type Configuration struct {
	errorHandling ErrorHandlingConfiguration
}

func (c Configuration) ErrorHandling() ErrorHandlingConfiguration {
	return c.errorHandling
}

func (c Configuration) Clone(opts ...ConfigurationOption) Configuration {
	newConfig := c
	for _, opt := range opts {
		opt(&newConfig)
	}
	return newConfig
}
