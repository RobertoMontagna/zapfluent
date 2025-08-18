package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestAsFluent(t *testing.T) {
	t.Run("with FluentEncoder", func(t *testing.T) {
		cfg := config.NewConfiguration()
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

		fluent := zapfluent.AsFluent(fluentEncoder)

		assert.NotNil(t, fluent)
	})

	t.Run("with other encoder", func(t *testing.T) {
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

		fluent := zapfluent.AsFluent(enc)

		assert.NotNil(t, fluent)
	})
}
