package config_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/config"
	"go.uber.org/zap/zapcore"
)

const (
	testFieldName    = "test-field"
	testErrorMessage = "test-error"
)

func TestFixedStringFallback(t *testing.T) {
	const fallbackValue = "fixed-value"
	factory := config.FixedStringFallback(fallbackValue)

	field := factory(testFieldName, errors.New(testErrorMessage))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, fallbackValue, enc.Fields[testFieldName])
}

func TestErrorStringFallback(t *testing.T) {
	const errorMsg = "this is the error message"
	factory := config.ErrorStringFallback()

	field := factory(testFieldName, errors.New(errorMsg))

	enc := zapcore.NewMapObjectEncoder()
	err := field.Encode(enc)
	assert.NoError(t, err)

	assert.Equal(t, errorMsg, enc.Fields[testFieldName])
}
