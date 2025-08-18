package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/config"
)

type FluentEncoder struct {
	config config.Configuration
	zapcore.Encoder
}

func NewFluentEncoder(
	encoder zapcore.Encoder,
	config config.Configuration,
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
