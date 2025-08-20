package zapfluent

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/config"
)

const (
	testError1 = "error 1"
	testError2 = "error 2"
)

func TestErrorHandler_Continue(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeContinue))
	handler := newErrorHandler(cfg)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.shouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.aggregatedError()

	assert.False(t, skip)
	assert.ErrorContains(t, finalErr, testError1)
	assert.ErrorContains(t, finalErr, testError2)
}

func TestErrorHandler_EarlyFailing(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeEarlyFailing))
	handler := newErrorHandler(cfg)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.shouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.aggregatedError()

	assert.True(t, skip)
	assert.ErrorContains(t, finalErr, testError1)
	assert.ErrorContains(t, finalErr, testError2)
}
