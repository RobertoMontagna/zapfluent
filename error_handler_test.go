package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
)

func TestErrorHandler_Continue(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeContinue))
	handler := zapfluent.NewErrorHandler(cfg)
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.ManageError(err1)
	skip := handler.ShouldSkip()
	handler.ManageError(err2)
	finalErr := handler.AggregatedError()

	assert.False(t, skip)
	assert.ErrorContains(t, finalErr, "error 1")
	assert.ErrorContains(t, finalErr, "error 2")
}

func TestErrorHandler_EarlyFailing(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeEarlyFailing))
	handler := zapfluent.NewErrorHandler(cfg)
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")

	handler.ManageError(err1)
	skip := handler.ShouldSkip()
	handler.ManageError(err2)
	finalErr := handler.AggregatedError()

	assert.True(t, skip)
	assert.Equal(t, err1, finalErr)
}
