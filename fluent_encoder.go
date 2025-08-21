package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// A FluentEncoder is a zapcore.Encoder that is aware of the zapfluent
// configuration. It wraps a standard zapcore.Encoder and is used to pass
// configuration details implicitly.
type FluentEncoder struct {
	config core.Configuration
	zapcore.Encoder
}

// NewFluentEncoder creates a new FluentEncoder that wraps the given
// zapcore.Encoder and holds the provided configuration.
func NewFluentEncoder(
	encoder zapcore.Encoder,
	config core.Configuration,
) *FluentEncoder {
	return &FluentEncoder{
		Encoder: encoder,
		config:  config,
	}
}

// Clone creates a copy of the FluentEncoder, including a clone of the
// underlying zapcore.Encoder and the associated configuration.
func (e *FluentEncoder) Clone() zapcore.Encoder {
	return &FluentEncoder{
		Encoder: e.Encoder.Clone(),
		config:  e.config.Clone(),
	}
}
