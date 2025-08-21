package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// AsFluent returns a new Fluent instance from a zapcore.ObjectEncoder.
//
// If the provided encoder is a `*FluentEncoder`, it uses the encoder's
// existing configuration. Otherwise, it creates a new Fluent instance with a
// default configuration. This is useful for integrating with libraries like
// Zap that provide an encoder.
func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	if fEnc, ok := encoder.(*FluentEncoder); ok {
		return NewFluent(fEnc, fEnc.config)
	}
	return NewFluent(encoder, core.NewConfiguration())
}
