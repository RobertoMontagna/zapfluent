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
	// Arrange
	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	// Act
	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	// Assert
	assert.NotNil(t, fluentEncoder)
}

func TestFluentEncoder_Clone(t *testing.T) {
	// Arrange
	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	// Act
	clone := fluentEncoder.Clone()

	// Assert
	assert.NotNil(t, clone)
	assert.NotSame(t, fluentEncoder, clone)
}
