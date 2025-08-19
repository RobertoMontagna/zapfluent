package zapfluent_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/testutil"
)

func TestErrorHandler_Continue(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeContinue))
	handler := zapfluent.NewErrorHandler(cfg)
	err1 := errors.New(testutil.TestError1)
	err2 := errors.New(testutil.TestError2)

	handler.AggregateError(err1)
	skip := handler.ShouldSkip()
	handler.AggregateError(err2)
	finalErr := handler.AggregatedError()

	assert.False(t, skip)
	assert.ErrorContains(t, finalErr, testutil.TestError1)
	assert.ErrorContains(t, finalErr, testutil.TestError2)
}

func TestErrorHandler_EarlyFailing(t *testing.T) {
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeEarlyFailing))
	handler := zapfluent.NewErrorHandler(cfg)
	err1 := errors.New(testutil.TestError1)
	err2 := errors.New(testutil.TestError2)

	handler.AggregateError(err1)
	skip := handler.ShouldSkip()
	handler.AggregateError(err2)
	finalErr := handler.AggregatedError()

	assert.True(t, skip)
	assert.ErrorContains(t, finalErr, testutil.TestError1)
	assert.ErrorContains(t, finalErr, testutil.TestError2)
}
