package zapfluent_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"

	. "github.com/onsi/gomega"
)

func TestAsFluent(t *testing.T) {
	g := NewWithT(t)

	t.Run("with FluentEncoder", func(t *testing.T) {
		cfg := zapfluent.NewConfiguration()
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
