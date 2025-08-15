package zapfluent

import "go.uber.org/zap/zapcore"

type Fluent struct {
	enc zapcore.ObjectEncoder
}

func NewFluent(
	enc zapcore.ObjectEncoder,
) *Fluent {
	fluent := &Fluent{
		enc: enc,
	}
	return fluent
}

func (z *Fluent) Add(encoder zapcore.ObjectMarshalerFunc) *Fluent {
	// TODO manage the error
	encoder(z.enc)
	return z
}

func (z *Fluent) Done() error {
	return nil
}
