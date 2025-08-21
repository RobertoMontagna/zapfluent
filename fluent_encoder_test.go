package zapfluent_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"

	. "github.com/onsi/gomega"
)

func TestNewFluentEncoder(t *testing.T) {
	g := NewWithT(t)

	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	g.Expect(fluentEncoder).ToNot(BeNil())
}

func TestFluentEncoder_Clone(t *testing.T) {
	g := NewWithT(t)

	cfg := config.NewConfiguration()
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	fluentEncoder := zapfluent.NewFluentEncoder(enc, cfg)

	clone := fluentEncoder.Clone()

	g.Expect(clone).ToNot(BeNil())
	g.Expect(clone).ToNot(BeIdenticalTo(fluentEncoder))
}
