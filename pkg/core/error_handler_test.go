package core

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

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

func TestErrorHandler_HandleError_ReturnsEmptyOptionalForNilError(t *testing.T) {
	g := NewWithT(t)
	cfg := NewErrorHandlingConfiguration()
	handler := NewErrorHandler(&cfg, nil)

	fallbackField := handler.HandleError(&testField{name: "test"}, nil)

	g.Expect(fallbackField.IsPresent()).To(BeFalse())
	g.Expect(handler.AggregatedError()).ToNot(HaveOccurred())
}

func TestErrorHandler_HandleError_AggregatesErrorWithoutFallback(t *testing.T) {
	g := NewWithT(t)
	cfg := NewErrorHandlingConfiguration()
	handler := NewErrorHandler(&cfg, nil)
	err := errors.New("test error")

	fallbackField := handler.HandleError(&testField{name: "test"}, err)

	g.Expect(fallbackField.IsPresent()).To(BeFalse())
	g.Expect(handler.AggregatedError()).To(MatchError(err))
}

func TestErrorHandler_HandleError_UsesFallbackFactory(t *testing.T) {
	g := NewWithT(t)
	fallbackValue := "fallback"
	cfg := NewErrorHandlingConfiguration(WithFallbackFieldFactory(FixedStringFallback(fallbackValue)))
	handler := NewErrorHandler(&cfg, nil)
	err := errors.New("test error")
	enc := zapcore.NewMapObjectEncoder()

	fallbackFieldOpt := handler.HandleError(&testField{name: "test"}, err)

	g.Expect(fallbackFieldOpt.IsPresent()).To(BeTrue())
	g.Expect(handler.AggregatedError()).To(MatchError(err))

	// Ensure the fallback field encodes correctly
	fallbackField, ok := fallbackFieldOpt.Get()
	g.Expect(ok).To(BeTrue())
	err = fallbackField.Encode(enc)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", fallbackValue))
}

func TestErrorHandler_EncodeField_Success(t *testing.T) {
	g := NewWithT(t)
	cfg := NewErrorHandlingConfiguration()
	enc := zapcore.NewMapObjectEncoder()
	handler := NewErrorHandler(&cfg, enc)
	field := String("test", "value")

	manager := handler.EncodeField(field)
	manager() // Execute the deferred encoding

	g.Expect(handler.AggregatedError()).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "value"))
}

func TestErrorHandler_EncodeField_FallbackSuccess(t *testing.T) {
	g := NewWithT(t)
	cfg := NewErrorHandlingConfiguration(WithFallbackFieldFactory(FixedStringFallback("fallback")))
	enc := zapcore.NewMapObjectEncoder()
	handler := NewErrorHandler(&cfg, enc)
	field := &testField{name: "test", encodeErr: errors.New("encode error")}

	manager := handler.EncodeField(field)
	manager()

	g.Expect(handler.AggregatedError()).To(HaveOccurred())
	g.Expect(handler.AggregatedError().Error()).To(ContainSubstring("encode error"))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "fallback"))
}

func TestErrorHandler_EncodeField_FallbackFails(t *testing.T) {
	g := NewWithT(t)
	fallbackErr := errors.New("fallback encode error")
	fallbackFactory := func(name string, err error) Field {
		return &testField{name: name, encodeErr: fallbackErr}
	}

	cfg := NewErrorHandlingConfiguration(WithFallbackFieldFactory(fallbackFactory))
	enc := zapcore.NewMapObjectEncoder()
	handler := NewErrorHandler(&cfg, enc)
	field := &testField{name: "test", encodeErr: errors.New("initial encode error")}

	manager := handler.EncodeField(field)
	manager()

	g.Expect(handler.AggregatedError()).To(HaveOccurred())
	g.Expect(handler.AggregatedError().Error()).To(ContainSubstring("initial encode error"))
	g.Expect(handler.AggregatedError().Error()).To(ContainSubstring("fallback encode error"))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "failed to encode fallback field"))
}

func TestErrorHandler_EncodeField_EarlyFailingSkip(t *testing.T) {
	g := NewWithT(t)
	cfg := NewErrorHandlingConfiguration(WithMode(ErrorHandlingModeEarlyFailing))
	enc := zapcore.NewMapObjectEncoder()
	handler := NewErrorHandler(&cfg, enc)

	// First, introduce an error
	handler.aggregateError(errors.New("first error"))

	// Now, try to encode a field
	field := String("test", "value")
	manager := handler.EncodeField(field)
	manager()

	g.Expect(enc.Fields).To(BeEmpty())
	g.Expect(handler.AggregatedError().Error()).To(Equal("first error"))
}

// testField is a mock implementation of the Field interface for testing purposes.
type testField struct {
	name      string
	encodeErr error
}

func (f *testField) Name() string {
	return f.name
}

func (f *testField) Encode(enc zapcore.ObjectEncoder) error {
	if f.encodeErr != nil {
		return f.encodeErr
	}
	enc.AddString(f.name, "dummy_value")
	return nil
}
