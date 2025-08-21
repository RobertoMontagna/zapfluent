package zapfluent_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent"

	. "github.com/onsi/gomega"
)

func TestNewConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := zapfluent.NewConfiguration()

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(zapfluent.ErrorHandlingModeContinue))
	})

	t.Run("with WithErrorHandling option", func(t *testing.T) {
		opt := zapfluent.WithErrorHandling(
			zapfluent.NewErrorHandlingConfiguration(
				zapfluent.WithMode(zapfluent.ErrorHandlingModeEarlyFailing),
			),
		)

		cfg := zapfluent.NewConfiguration(opt)

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(zapfluent.ErrorHandlingModeEarlyFailing))
	})
}

func TestConfiguration_Clone(t *testing.T) {
	g := NewWithT(t)
	originalCfg := zapfluent.NewConfiguration(
		zapfluent.WithErrorHandling(
			zapfluent.NewErrorHandlingConfiguration(
				zapfluent.WithMode(zapfluent.ErrorHandlingModeEarlyFailing),
			),
		),
	)

	clone := originalCfg.Clone()

	g.Expect(clone).To(Equal(originalCfg))
}
