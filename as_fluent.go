package zapfluent

import "go.uber.org/zap/zapcore"

func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fEnc, ok := encoder.(*FluentEncoder); ok {
		return NewFluent(fEnc, fEnc.config)
	}
	return NewFluent(encoder, NewFluentConfig())
}
