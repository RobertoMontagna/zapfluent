package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestErrorHandlingMode_String(t *testing.T) {
	testCases := []struct {
		mode     config.ErrorHandlingMode
		expected string
	}{
		{config.ErrorHandlingModeUnknown, config.ErrorHandlingModeUnknownString},
		{config.ErrorHandlingModeEarlyFailing, config.ErrorHandlingModeEarlyFailingString},
		{config.ErrorHandlingModeContinue, config.ErrorHandlingModeContinueString},
		{config.ErrorHandlingMode(99), "Unknown(99)"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			s := tc.mode.String()

			assert.Equal(t, tc.expected, s)
		})
	}
}

func TestIntToErrorHandlingMode(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected config.ErrorHandlingMode
	}{
		{config.ErrorHandlingModeUnknownString, 0, config.ErrorHandlingModeUnknown},
		{config.ErrorHandlingModeEarlyFailingString, 1, config.ErrorHandlingModeEarlyFailing},
		{config.ErrorHandlingModeContinueString, 2, config.ErrorHandlingModeContinue},
		{"Invalid", 99, config.ErrorHandlingModeUnknown},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mode := config.IntToErrorHandlingMode(tc.value)

			assert.Equal(t, tc.expected, mode)
		})
	}
}
