package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

// Fluent is a wrapper around a zapcore.ObjectEncoder that provides a fluent
// interface for adding fields. It is designed to simplify the process of
// encoding structured, nested data.
type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *ErrorHandler
}

// NewFluent creates a new Fluent instance.
//
// It requires a zapcore.ObjectEncoder for encoding fields and a
// Configuration to define its behavior, particularly for error handling.
func NewFluent(
	enc zapcore.ObjectEncoder,
	config config.Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: NewErrorHandler(config.ErrorHandling()),
	}
}

// Add encodes the given Field using the underlying zapcore.ObjectEncoder.
//
// If the field fails to encode, Add handles the error according to the configured
// error handling strategy. This may include stopping further processing,
// using a fallback field, or aggregating the error.
//
// It returns the Fluent instance to allow for method chaining.
func (f *Fluent) Add(field fluentfield.Field) *Fluent {
	if f.errorHandler.ShouldSkip() {
		return f
	}

	f.errorHandler.Process(field, field.Encode(f.enc)).ForEach(func(fallbackField fluentfield.Field) {
		if err := fallbackField.Encode(f.enc); err != nil {
			f.errorHandler.AggregateError(err)
		}
	})

	return f
}

// Done returns the final aggregated error after all fields have been processed.
//
// If no errors occurred during the encoding of any fields, it returns nil.
// Otherwise, it returns an error that may contain multiple underlying errors,
// which can be inspected using the `multierr` package.
func (f *Fluent) Done() error {
	return f.errorHandler.AggregatedError()
}
