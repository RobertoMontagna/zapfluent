// Package zapfluent provides a fluent interface for structured logging with Zap.
// It allows for a more intuitive and chainable way to add fields to a log entry.
package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

// Fluent provides a fluent interface for adding structured logging fields to a
// Zap ObjectEncoder. It is designed to be used in a chainable manner.
type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *errorHandler
}

// NewFluent creates and returns a new Fluent instance.
// It requires a Zap ObjectEncoder to write the fields to, and a Configuration
// object to configure its behavior, such as error handling.
func NewFluent(
	enc zapcore.ObjectEncoder,
	config config.Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: newErrorHandler(config.ErrorHandling()),
	}
}

// Add adds a field to the log entry.
// It takes a fluentfield.Field, which is an interface that allows for custom
// field types and encoding logic.
// The method returns the Fluent pointer, allowing for chained calls.
func (z *Fluent) Add(field fluentfield.Field) *Fluent {
	if z.errorHandler.shouldSkip() {
		return z
	}

	maybeFallbackField := z.errorHandler.handleError(field, field.Encode(z.enc))

	maybeEncodingError := optional.FlatMap(maybeFallbackField, z.encodeAndLift)

	optional.Map(maybeEncodingError, func(encodingErr error) bool {
		if encodingErr != nil {
			z.enc.AddString("fluent_error", "failed to encode fallback field")
		}
		return true
	})

	return z
}

func (z *Fluent) encodeAndLift(field fluentfield.Field) optional.Optional[error] {
	return optional.OfError(field.Encode(z.enc))
}

// Done completes the fluent chain and returns any aggregated errors that
// occurred during the process. This should be the final call in the chain.
func (z *Fluent) Done() error {
	return z.errorHandler.aggregatedError()
}
