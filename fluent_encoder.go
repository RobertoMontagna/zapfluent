package zapfluent

import (
	"go.uber.org/zap/zapcore"
)

type FluentEncoder struct {
	config FluentConfig
	zapcore.Encoder
}

func NewFluentEncoder(
	encoder zapcore.Encoder,
	config FluentConfig,
) *FluentEncoder {
	return &FluentEncoder{
		Encoder: encoder,
		config:  config,
	}
}

func (e *FluentEncoder) Clone() zapcore.Encoder {
	return &FluentEncoder{
		Encoder: e.Encoder.Clone(),
		config:  e.config.Clone(),
	}
}

