package zapfluent

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/config"
)

const (
	testError1 = "error 1"
	testError2 = "error 2"
)

func TestErrorHandler_Continue(t *testing.T) {
	g := NewWithT(t)
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeContinue))
	handler := newErrorHandler(cfg, nil)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.shouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.aggregatedError()

	g.Expect(skip).To(BeFalse())
	g.Expect(finalErr).To(HaveOccurred())
	g.Expect(finalErr.Error()).To(ContainSubstring(testError1))
	g.Expect(finalErr.Error()).To(ContainSubstring(testError2))
}

func TestErrorHandler_EarlyFailing(t *testing.T) {
	g := NewWithT(t)
	cfg := config.NewErrorHandlingConfiguration(config.WithMode(config.ErrorHandlingModeEarlyFailing))
	handler := newErrorHandler(cfg, nil)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.shouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.aggregatedError()

	g.Expect(skip).To(BeTrue())
	g.Expect(finalErr).To(HaveOccurred())
	g.Expect(finalErr.Error()).To(ContainSubstring(testError1))
	g.Expect(finalErr.Error()).To(ContainSubstring(testError2))
}
