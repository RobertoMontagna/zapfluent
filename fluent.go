package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

type Fluent struct {
	enc          zapcore.ObjectEncoder
	errorHandler *errorHandler
}

func NewFluent(
	enc zapcore.ObjectEncoder,
	config config.Configuration,
) *Fluent {
	return &Fluent{
		enc:          enc,
		errorHandler: newErrorHandler(config.ErrorHandling()),
	}
}

func (z *Fluent) Add(field fluentfield.Field) *Fluent {
	if z.errorHandler.shouldSkip() {
		return z
	}

	z.errorHandler.process(field, field.Encode(z.enc)).ForEach(func(fallbackField fluentfield.Field) {
		if err := fallbackField.Encode(z.enc); err != nil {
			z.errorHandler.aggregateError(err)
		}
	})

	return z
}

func (z *Fluent) Done() error {
	return z.errorHandler.aggregatedError()
}
