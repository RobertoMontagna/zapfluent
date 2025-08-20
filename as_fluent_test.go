package zapfluent_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestAsFluent(t *testing.T) {
	g := NewWithT(t)

	t.Run("with FluentEncoder", func(t *testing.T) {
		cfg := config.NewConfiguration()
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
		fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

		fluent := zapfluent.AsFluent(fluentEncoder)

		g.Expect(fluent).ToNot(BeNil())
	})

	t.Run("with other encoder", func(t *testing.T) {
		enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

		fluent := zapfluent.AsFluent(enc)

		g.Expect(fluent).ToNot(BeNil())
	})
}
