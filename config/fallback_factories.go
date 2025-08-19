package config

import "go.robertomontagna.dev/zapfluent/fluentfield"

// FixedStringFallback returns a factory that creates a string field with a fixed value.
// The original error is ignored.
func FixedStringFallback(value string) FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, value)
	}
}

// ErrorStringFallback returns a factory that creates a string field with the
// error message as its value.
func ErrorStringFallback() FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, err.Error())
	}
}
