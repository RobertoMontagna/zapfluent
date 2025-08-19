package config

import "go.robertomontagna.dev/zapfluent/fluentfield"

func FixedStringFallback(value string) FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, value)
	}
}

func ErrorStringFallback() FallbackFieldFactory {
	return func(name string, err error) fluentfield.Field {
		return fluentfield.String(name, err.Error())
	}
}
