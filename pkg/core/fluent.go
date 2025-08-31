package core

import (
	"go.uber.org/zap/zapcore"
)

// Fluent provides a fluent interface for adding structured logging fields to a
// Zap ObjectEncoder. It is designed to be used in a chainable manner.
type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *ErrorHandler
}

// NewFluent creates and returns a new Fluent instance.
// It requires a Zap ObjectEncoder to write the fields to, and a Configuration
// NewFluent creates a Fluent that writes structured fields to the provided zapcore.ObjectEncoder.
// The returned Fluent holds the encoder and initializes its error handler from the given Configuration's
// error-handling settings.
func NewFluent(
	enc zapcore.ObjectEncoder,
	config Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: NewErrorHandler(config.ErrorHandling(), enc),
	}
}

// Add adds a field to the log entry.
// It takes a core.Field, which is an interface that allows for custom
// field types and encoding logic.
// The method returns the Fluent pointer, allowing for chained calls.
func (z *Fluent) Add(field Field) *Fluent {
	if z.errorHandler.ShouldSkip() {
		return z
	}

	encodingErrorManager := z.errorHandler.EncodeField(field)
	encodingErrorManager()

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
// AsFluent converts a zapcore.ObjectEncoder into a *Fluent.
// If the provided encoder is already a *FluentEncoder, its Config is reused;
// otherwise a new Fluent is created using the default Configuration.
func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fEnc, ok := encoder.(*FluentEncoder); ok {
		return NewFluent(fEnc, fEnc.Config)
	}
	return NewFluent(encoder, NewConfiguration())
}
