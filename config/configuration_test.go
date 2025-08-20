package config_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/config"
)

func TestNewConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := config.NewConfiguration()

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(config.ErrorHandlingModeContinue))
	})

	t.Run("with WithErrorHandling option", func(t *testing.T) {
		opt := config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		)

		cfg := config.NewConfiguration(opt)

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(config.ErrorHandlingModeEarlyFailing))
	})
}

func TestConfiguration_Clone(t *testing.T) {
	g := NewWithT(t)
	originalCfg := config.NewConfiguration(
		config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		),
	)

	clone := originalCfg.Clone()

	g.Expect(clone).To(Equal(originalCfg))
}
