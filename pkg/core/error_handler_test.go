package core

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	testError1 = "error 1"
	testError2 = "error 2"
)

func TestErrorHandler_Continue(t *testing.T) {
	g := NewWithT(t)

	cfg := NewErrorHandlingConfiguration(WithMode(ErrorHandlingModeContinue))
	handler := NewErrorHandler(&cfg, nil)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.ShouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.AggregatedError()

	g.Expect(skip).To(BeFalse())
	g.Expect(finalErr).To(HaveOccurred())
	g.Expect(finalErr.Error()).To(ContainSubstring(testError1))
	g.Expect(finalErr.Error()).To(ContainSubstring(testError2))
}

func TestErrorHandler_EarlyFailing(t *testing.T) {
	g := NewWithT(t)

	cfg := NewErrorHandlingConfiguration(WithMode(ErrorHandlingModeEarlyFailing))
	handler := NewErrorHandler(&cfg, nil)
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)

	handler.aggregateError(err1)
	skip := handler.ShouldSkip()
	handler.aggregateError(err2)
	finalErr := handler.AggregatedError()

	g.Expect(skip).To(BeTrue())
	g.Expect(finalErr).To(HaveOccurred())
	g.Expect(finalErr.Error()).To(ContainSubstring(testError1))
	g.Expect(finalErr.Error()).To(ContainSubstring(testError2))
}
