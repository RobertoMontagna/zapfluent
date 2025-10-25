package core

// A ConfigurationOption is a function that applies a configuration to a
// Configuration object.
type ConfigurationOption func(*Configuration)

// WithErrorHandling is a ConfigurationOption that sets the error handling
// error handling configuration to the provided ErrorHandlingConfiguration.
func WithErrorHandling(errorHandling ErrorHandlingConfiguration) ConfigurationOption {
	return func(c *Configuration) {
		c.errorHandling = errorHandling
	}
}

// Configuration holds the settings for a Fluent instance.
type Configuration struct {
	errorHandling ErrorHandlingConfiguration
}

// NewConfiguration creates a new Configuration with the given options.
//
// NewConfiguration creates a Configuration initialized with a default error handling configuration.
// It applies each provided ConfigurationOption in order to the configuration; if no options are provided the default configuration is returned.
func NewConfiguration(opts ...ConfigurationOption) Configuration {
	config := Configuration{
		errorHandling: NewErrorHandlingConfiguration(),
	}
	for _, opt := range opts {
		opt(&config)
	}
	return config
}

// ErrorHandling returns the error handling configuration.
func (c *Configuration) ErrorHandling() *ErrorHandlingConfiguration {
	return &c.errorHandling
}

// Clone creates a shallow copy of the Configuration.
func (c *Configuration) Clone() Configuration {
	return *c
}