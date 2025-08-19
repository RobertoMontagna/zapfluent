package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *ErrorHandler
}

func NewFluent(
	enc zapcore.ObjectEncoder,
	config config.Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: NewErrorHandler(config.ErrorHandling()),
	}
}

func (z *Fluent) Add(field fluentfield.Field) *Fluent {
	if z.errorHandler.ShouldSkip() {
		return z
	}

	z.errorHandler.Process(field, field.Encode(z.enc)).ForEach(func(fallbackField fluentfield.Field) {
		if err := fallbackField.Encode(z.enc); err != nil {
			z.errorHandler.AggregateError(err)
		}
	})

	return z
}

func (z *Fluent) Done() error {
	return z.errorHandler.AggregatedError()
}
