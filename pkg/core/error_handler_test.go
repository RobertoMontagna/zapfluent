package core_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
	"go.robertomontagna.dev/zapfluent/testutil/stubs"

	. "github.com/onsi/gomega"
)

var (
	errTest1 = errors.New("error 1")
)

func TestErrorHandler_ShouldSkip_ContinueMode(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeContinue))
	handler := core.NewErrorHandler(&cfg, zapcore.NewMapObjectEncoder())

	handler.EncodeField(stubs.NewFailingField("test", errTest1))()

	g.Expect(handler.ShouldSkip()).To(BeFalse())
	g.Expect(handler.AggregatedError()).To(HaveOccurred())
}

func TestErrorHandler_ShouldSkip_EarlyFailingMode(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeEarlyFailing))
	handler := core.NewErrorHandler(&cfg, zapcore.NewMapObjectEncoder())

	handler.EncodeField(stubs.NewFailingField("test", errTest1))()

	g.Expect(handler.ShouldSkip()).To(BeTrue())
	g.Expect(handler.AggregatedError()).To(HaveOccurred())
}

func TestErrorHandler_HandleError_ReturnsEmptyOptionalForNilError(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration()
	handler := core.NewErrorHandler(&cfg, nil)

	fallbackField := handler.HandleError(core.String("test", "value"), nil)

	g.Expect(fallbackField.IsPresent()).To(BeFalse())
	g.Expect(handler.AggregatedError()).ToNot(HaveOccurred())
}

func TestErrorHandler_HandleError_AggregatesErrorWithoutFallback(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration()
	handler := core.NewErrorHandler(&cfg, nil)

	fallbackField := handler.HandleError(core.String("test", "value"), errTest1)

	g.Expect(fallbackField.IsPresent()).To(BeFalse())
	g.Expect(handler.AggregatedError()).To(MatchError(errTest1))
}

func TestErrorHandler_HandleError_UsesFallbackFactory(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithFallbackFieldFactory(core.FixedStringFallback("fallback")))
	handler := core.NewErrorHandler(&cfg, nil)
	enc := zapcore.NewMapObjectEncoder()

	fallbackFieldOpt := handler.HandleError(core.String("test", "value"), errTest1)

	g.Expect(fallbackFieldOpt.IsPresent()).To(BeTrue())
	g.Expect(handler.AggregatedError()).To(MatchError(errTest1))

	fallbackField, ok := fallbackFieldOpt.Get()
	g.Expect(ok).To(BeTrue())
	g.Expect(fallbackField.Encode(enc)).To(Succeed())
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "fallback"))
}

func TestErrorHandler_EncodeField_Success(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration()
	handler := core.NewErrorHandler(&cfg, zapcore.NewMapObjectEncoder())

	handler.EncodeField(core.String("test", "value"))()

	g.Expect(handler.AggregatedError()).ToNot(HaveOccurred())
}

func TestErrorHandler_EncodeField_FallbackSuccess(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithFallbackFieldFactory(core.FixedStringFallback("fallback")))
	enc := zapcore.NewMapObjectEncoder()
	handler := core.NewErrorHandler(&cfg, enc)
	errEncode := errors.New("encode error")

	handler.EncodeField(stubs.NewFailingField("test", errEncode))()

	g.Expect(handler.AggregatedError()).To(MatchError(errEncode))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "fallback"))
}

func TestErrorHandler_EncodeField_FallbackFails(t *testing.T) {
	g := NewWithT(t)
	errInitial := errors.New("initial encode error")
	errFallback := errors.New("fallback encode error")
	fallbackFactory := func(name string, err error) core.Field {
		return stubs.NewFailingField(name, errFallback)
	}
	cfg := core.NewErrorHandlingConfiguration(core.WithFallbackFieldFactory(fallbackFactory))
	enc := zapcore.NewMapObjectEncoder()
	handler := core.NewErrorHandler(&cfg, enc)

	handler.EncodeField(stubs.NewFailingField("test", errInitial))()

	g.Expect(handler.AggregatedError()).To(HaveOccurred())
	g.Expect(handler.AggregatedError()).To(MatchError(errInitial))
	g.Expect(handler.AggregatedError()).To(MatchError(errFallback))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "failed to encode fallback field"))
}

func TestErrorHandler_EncodeField_EarlyFailingSkip(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeEarlyFailing))
	enc := zapcore.NewMapObjectEncoder()
	handler := core.NewErrorHandler(&cfg, enc)

	handler.EncodeField(stubs.NewFailingField("first", errTest1))()
	handler.EncodeField(core.String("second", "value"))()

	g.Expect(handler.AggregatedError()).To(MatchError(errTest1))
	g.Expect(enc.Fields).To(BeEmpty())
}
