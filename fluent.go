// Package zapfluent provides a fluent interface for structured logging with Zap.
// It allows for a more intuitive and chainable way to add fields to a log entry.
package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// Fluent provides a fluent interface for adding structured logging fields to a
// Zap ObjectEncoder. It is designed to be used in a chainable manner.
type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *core.ErrorHandler
}

// NewFluent creates and returns a new Fluent instance.
// It requires a Zap ObjectEncoder to write the fields to, and a Configuration
// object to configure its behavior, such as error handling.
func NewFluent(
	enc zapcore.ObjectEncoder,
	config core.Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: core.NewErrorHandler(config.ErrorHandling(), enc),
	}
}

// Add adds a field to the log entry.
// It takes a core.Field, which is an interface that allows for custom
// field types and encoding logic.
// The method returns the Fluent pointer, allowing for chained calls.
func (z *Fluent) Add(field core.Field) *Fluent {
	if z.errorHandler.ShouldSkip() {
		return z
	}

	z.errorHandler.EncodeField(field)()

	return z
}

// Done completes the fluent chain and returns any aggregated errors that
// occurred during the process. This should be the final call in the chain.
func (z *Fluent) Done() error {
	return z.errorHandler.AggregatedError()
}

// AsFluent returns a new Fluent instance from a zapcore.ObjectEncoder.
//
// If the provided encoder is a `*FluentEncoder`, it uses the encoder's
// existing configuration. Otherwise, it creates a new Fluent instance with a
// default configuration. This is useful for integrating with libraries like
// Zap that provide an encoder.
func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fEnc, ok := encoder.(*core.FluentEncoder); ok {
		return NewFluent(fEnc, fEnc.Config)
	}
	return NewFluent(encoder, core.NewConfiguration())
}
