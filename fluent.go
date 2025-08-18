package zapfluent

import (
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/fluentfield"
)

type Fluent struct {
	enc    zapcore.ObjectEncoder
	err    error
	config FluentConfig
}

func NewFluent(
	enc zapcore.ObjectEncoder,
	config FluentConfig,
) *Fluent {
	fluent := &Fluent{
		enc:    enc,
		config: config,
	}
	return fluent
}

func (z *Fluent) Add(field fluentfield.Field) *Fluent {
	if z.err != nil {
		return z
	}
	if err := field.Encode(z.enc); err != nil {
		z.err = multierr.Append(z.err, err)
	}
	return z
}

func (z *Fluent) Done() error {
	return z.err
}
