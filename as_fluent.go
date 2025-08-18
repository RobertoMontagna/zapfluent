package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
)

func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fluentEncoder, isFluentEncoder := encoder.(*FluentEncoder); isFluentEncoder {
		return NewFluent(fluentEncoder, fluentEncoder.config)
	}
	return NewFluent(encoder, config.NewConfiguration())
}
