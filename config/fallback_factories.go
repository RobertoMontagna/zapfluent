package config

import "go.robertomontagna.dev/zapfluent/fluentfield"

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
