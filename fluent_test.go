package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent"
	"go.uber.org/zap/zapcore"
)

type errorField struct {
	err error
}

func (f errorField) Encode(enc zapcore.ObjectEncoder) error {
	return f.err
}

func TestFluent_errorHandling(t *testing.T) {
	expectedErr := errors.New("test error")

	fluent := zapfluent.NewFluent(nil)
	fluent.Add(errorField{err: expectedErr})
	err := fluent.Done()

	assert.Equal(t, expectedErr, err)
}
