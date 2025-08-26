package core_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

func TestNewConfiguration(t *testing.T) {
	testCases := []struct {
		name         string
		options      []core.ConfigurationOption
		expectedMode core.ErrorHandlingMode
	}{
		{
			name:         "with default options",
			options:      []core.ConfigurationOption{},
			expectedMode: core.ErrorHandlingModeContinue,
		},
		{
			name: "with WithErrorHandling option",
			options: []core.ConfigurationOption{
				core.WithErrorHandling(
					core.NewErrorHandlingConfiguration(
						core.WithMode(core.ErrorHandlingModeEarlyFailing),
					),
				),
			},
			expectedMode: core.ErrorHandlingModeEarlyFailing,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			cfg := core.NewConfiguration(tc.options...)

			g.Expect(cfg.ErrorHandling().Mode()).To(Equal(tc.expectedMode))
		})
	}
}

func TestConfiguration_Clone(t *testing.T) {
	g := NewWithT(t)

	originalCfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithMode(core.ErrorHandlingModeEarlyFailing),
			),
		),
	)

	clone := originalCfg.Clone()

	g.Expect(clone).To(Equal(originalCfg))
}
