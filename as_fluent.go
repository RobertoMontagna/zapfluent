package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
)

func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fEnc, ok := encoder.(*FluentEncoder); ok {
		return NewFluent(fEnc, fEnc.config)
	}
	return NewFluent(encoder, config.NewConfiguration())
}
