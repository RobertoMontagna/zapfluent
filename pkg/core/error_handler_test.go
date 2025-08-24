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
	errTest1                     = errors.New("error 1")
	errEncode                    = errors.New("encode error")
	errInitial                   = errors.New("initial encode error")
	failingField                 = stubs.NewFailingFieldForTest(stubs.WithName("test"), stubs.WithError(errTest1))
	anotherFailingField          = stubs.NewFailingFieldForTest(stubs.WithName("first"), stubs.WithError(errTest1))
	failingFieldWithEncodeError  = stubs.NewFailingFieldForTest(stubs.WithName("test"), stubs.WithError(errEncode))
	failingFieldWithInitialError = stubs.NewFailingFieldForTest(stubs.WithName("test"), stubs.WithError(errInitial))
)

func TestErrorHandler_ShouldSkip_ContinueMode(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeContinue))
	handler := core.NewErrorHandler(&cfg, zapcore.NewMapObjectEncoder())

	handler.EncodeField(failingField)()

	g.Expect(handler.ShouldSkip()).To(BeFalse())
	g.Expect(handler.AggregatedError()).To(HaveOccurred())
}

func TestErrorHandler_ShouldSkip_EarlyFailingMode(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeEarlyFailing))
	handler := core.NewErrorHandler(&cfg, zapcore.NewMapObjectEncoder())

	handler.EncodeField(failingField)()

	g.Expect(handler.ShouldSkip()).To(BeTrue())
	g.Expect(handler.AggregatedError()).To(HaveOccurred())
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

	handler.EncodeField(failingFieldWithEncodeError)()

	g.Expect(handler.AggregatedError()).To(MatchError(errEncode))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "fallback"))
}

func TestErrorHandler_EncodeField_FallbackFails(t *testing.T) {
	g := NewWithT(t)
	errFallback := errors.New("fallback encode error")
	fallbackFactory := func(name string, err error) core.Field {
		return stubs.NewFailingFieldForTest(
			stubs.WithName(name),
			stubs.WithError(errFallback),
		)
	}
	cfg := core.NewErrorHandlingConfiguration(core.WithFallbackFieldFactory(fallbackFactory))
	enc := zapcore.NewMapObjectEncoder()
	handler := core.NewErrorHandler(&cfg, enc)

	handler.EncodeField(failingFieldWithInitialError)()

	g.Expect(handler.AggregatedError()).To(MatchError(errInitial))
	g.Expect(handler.AggregatedError()).To(MatchError(errFallback))
	g.Expect(enc.Fields).To(HaveKeyWithValue("test", "failed to encode fallback field"))
}

func TestErrorHandler_EncodeField_EarlyFailingSkip(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewErrorHandlingConfiguration(core.WithMode(core.ErrorHandlingModeEarlyFailing))
	enc := zapcore.NewMapObjectEncoder()
	handler := core.NewErrorHandler(&cfg, enc)

	handler.EncodeField(anotherFailingField)()
	handler.EncodeField(core.String("second", "value"))()

	g.Expect(handler.AggregatedError()).To(MatchError(errTest1))
	g.Expect(enc.Fields).To(BeEmpty())
}
