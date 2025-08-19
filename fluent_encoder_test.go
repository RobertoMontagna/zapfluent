package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestNewFluentEncoder(t *testing.T) {
	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	assert.NotNil(t, fluentEncoder)
}

func TestFluentEncoder_Clone(t *testing.T) {
	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	clone := fluentEncoder.Clone()

	assert.NotNil(t, clone)
	assert.NotSame(t, fluentEncoder, clone)
}
