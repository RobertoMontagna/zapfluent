package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestErrorHandlingMode_String(t *testing.T) {
	// Arrange
	testCases := []struct {
		mode     config.ErrorHandlingMode
		expected string
	}{
		{config.ErrorHandlingModeUnknown, "Unknown"},
		{config.ErrorHandlingModeEarlyFailing, "EarlyFailing"},
		{config.ErrorHandlingModeContinue, "Continue"},
		{config.ErrorHandlingMode(99), "Unknown(99)"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			// Act
			s := tc.mode.String()

			// Assert
			assert.Equal(t, tc.expected, s)
		})
	}
}

func TestIntToErrorHandlingMode(t *testing.T) {
	// Arrange
	testCases := []struct {
		name     string
		value    int
		expected config.ErrorHandlingMode
	}{
		{"Unknown", 0, config.ErrorHandlingModeUnknown},
		{"EarlyFailing", 1, config.ErrorHandlingModeEarlyFailing},
		{"Continue", 2, config.ErrorHandlingModeContinue},
		{"Invalid", 99, config.ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			mode := config.IntToErrorHandlingMode(tc.value)

			// Assert
			assert.Equal(t, tc.expected, mode)
		})
	}
}
