package config

// A ConfigurationOption is a function that applies a configuration to a
// Configuration object.
type ConfigurationOption func(*Configuration)

// WithErrorHandling is a ConfigurationOption that sets the error handling
// strategy for the Fluent instance.
func WithErrorHandling(errorHandling ErrorHandlingConfiguration) ConfigurationOption {
	return func(c *Configuration) {
		c.errorHandling = errorHandling
	}
}

// NewConfiguration creates a new Configuration with the given options.
//
// If no options are provided, it returns a default configuration.
func NewConfiguration(opts ...ConfigurationOption) Configuration {
	config := Configuration{
		errorHandling: NewErrorHandlingConfiguration(),
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

// Configuration holds the settings for a Fluent instance.
type Configuration struct {
	errorHandling ErrorHandlingConfiguration
}

// ErrorHandling returns the error handling configuration.
func (c Configuration) ErrorHandling() ErrorHandlingConfiguration {
	return c.errorHandling
}

// Clone creates a shallow copy of the Configuration.
func (c Configuration) Clone() Configuration {
	return c
}
