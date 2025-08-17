package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

type errorField struct {
	err error
}

func (f errorField) Encode(enc zapcore.ObjectEncoder) error {
	return f.err
}

func TestFluent_errorHandling(t *testing.T) {
	expectedErr := errors.New("test error")
	field := errorField{err: expectedErr}
	fluent := zapfluent.NewFluent(nil)

	err := fluent.Add(field).Done()

	assert.Equal(t, expectedErr, err)
}
