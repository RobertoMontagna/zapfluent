package zapfluent

import (
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
)

type Fluent struct {
	enc zapcore.ObjectEncoder
	err error
}

func NewFluent(
	enc zapcore.ObjectEncoder,
) *Fluent {
	fluent := &Fluent{
		enc: enc,
	}
	return fluent
}

func (z *Fluent) Add(field Field) *Fluent {
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
