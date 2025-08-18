package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

func TestAsFluent(t *testing.T) {
	t.Run("with FluentEncoder", func(t *testing.T) {
		cfg := config.NewConfiguration()
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

		fluent := zapfluent.AsFluent(fluentEncoder)

		assert.NotNil(t, fluent)
		err := fluent.Add(fluentfield.String("a-key", "a-value")).Done()
		assert.NoError(t, err)
	})

	t.Run("with other encoder", func(t *testing.T) {
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

		fluent := zapfluent.AsFluent(enc)

		assert.NotNil(t, fluent)
		err := fluent.Add(fluentfield.String("a-key", "a-value")).Done()
		assert.NoError(t, err)
	})
}
