package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
)

type errorField struct {
	err error
}

func (f errorField) Encode(enc zapcore.ObjectEncoder) error {
	return f.err
}

func (f errorField) Name() string {
	return "error"
}

func TestFluent_errorHandling(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := errorField{err: err1}
	field2 := errorField{err: err2}
	fluent := zapfluent.AsFluent(nil)

	err := fluent.Add(field1).Add(field2).Done()

	assert.ErrorContains(t, err, "error 1")
	assert.ErrorContains(t, err, "error 2")
}

func TestFluent_errorHandling_EarlyFailing(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	field1 := errorField{err: err1}
	field2 := errorField{err: err2}

	cfg := config.NewConfiguration(config.WithErrorHandling(config.WithMode(config.ErrorHandlingModeEarlyFailing)))

	// Create a dummy encoder
	enc := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

	// Create a fluent instance with the custom config
	fluent := zapfluent.NewFluent(enc, cfg)

	err := fluent.Add(field1).Add(field2).Done()

	assert.Equal(t, err1, err) // Should only return the first error
}
