package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/config"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("with default options", func(t *testing.T) {
		cfg := config.NewConfiguration()

		assert.Equal(t, config.ErrorHandlingModeContinue, cfg.ErrorHandling().Mode())
	})

	t.Run("with WithErrorHandling option", func(t *testing.T) {
		opt := config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		)

		cfg := config.NewConfiguration(opt)

		assert.Equal(t, config.ErrorHandlingModeEarlyFailing, cfg.ErrorHandling().Mode())
	})
}

func TestConfiguration_Clone(t *testing.T) {
	originalCfg := config.NewConfiguration(
		config.WithErrorHandling(
			config.NewErrorHandlingConfiguration(
				config.WithMode(config.ErrorHandlingModeEarlyFailing),
			),
		),
	)

	clone := originalCfg.Clone()

	assert.Equal(t, originalCfg, clone)
}
