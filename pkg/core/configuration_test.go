package core

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewConfiguration(t *testing.T) {
	g := NewWithT(t)

	t.Run("with default options", func(t *testing.T) {
		cfg := NewConfiguration()

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(ErrorHandlingModeContinue))
	})

	t.Run("with WithErrorHandling option", func(t *testing.T) {
		opt := WithErrorHandling(
			NewErrorHandlingConfiguration(
				WithMode(ErrorHandlingModeEarlyFailing),
			),
		)

		cfg := NewConfiguration(opt)

		g.Expect(cfg.ErrorHandling().Mode()).To(Equal(ErrorHandlingModeEarlyFailing))
	})
}

func TestConfiguration_Clone(t *testing.T) {
	g := NewWithT(t)
	originalCfg := NewConfiguration(
		WithErrorHandling(
			NewErrorHandlingConfiguration(
				WithMode(ErrorHandlingModeEarlyFailing),
			),
		),
	)

	clone := originalCfg.Clone()

	g.Expect(clone).To(Equal(originalCfg))
}
