package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/config"
)

func TestNewErrorHandlingConfiguration(t *testing.T) {
	t.Run("with default options", func(t *testing.T) {
		cfg := config.NewErrorHandlingConfiguration()

		assert.Equal(t, config.ErrorHandlingModeContinue, cfg.Mode())
	})

	t.Run("with WithMode option", func(t *testing.T) {
		opt := config.WithMode(config.ErrorHandlingModeEarlyFailing)

		cfg := config.NewErrorHandlingConfiguration(opt)

		assert.Equal(t, config.ErrorHandlingModeEarlyFailing, cfg.Mode())
	})
}
