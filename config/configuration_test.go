package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("with default options", func(t *testing.T) {
		// Arrange & Act
		cfg := config.NewConfiguration()

		// Assert
		assert.Equal(t, config.ErrorHandlingModeContinue, cfg.ErrorHandling().Mode())
	})

	t.Run("with WithErrorHandling option", func(t *testing.T) {
		// Arrange
		opt := config.WithErrorHandling(config.WithMode(config.ErrorHandlingModeEarlyFailing))

		// Act
		cfg := config.NewConfiguration(opt)

		// Assert
		assert.Equal(t, config.ErrorHandlingModeEarlyFailing, cfg.ErrorHandling().Mode())
	})
}

func TestConfiguration_Clone(t *testing.T) {
	// Arrange
	originalCfg := config.NewConfiguration(
		config.WithErrorHandling(config.WithMode(config.ErrorHandlingModeEarlyFailing)),
	)

	// Act
	clone := originalCfg.Clone()

	// Assert
	assert.Equal(t, originalCfg, clone)
}
