package config_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/config"
	"go.uber.org/zap/zapcore"
)

func TestFixedStringFallback(t *testing.T) {
	const fallbackValue = "fixed-value"
	factory := config.FixedStringFallback(fallbackValue)

	field := factory("test-field", errors.New("test-error"))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, fallbackValue, enc.Fields["test-field"])
}

func TestErrorStringFallback(t *testing.T) {
	const errorMsg = "this is the error message"
	factory := config.ErrorStringFallback()

	field := factory("test-field", errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, errorMsg, enc.Fields["test-field"])
}
