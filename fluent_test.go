package zapfluent_test

import (
	"errors"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/pkg/core"
	"go.robertomontagna.dev/zapfluent/testutil/stubs"

	. "github.com/onsi/gomega"
)

var (
	errTest1          = errors.New("error 1")
	errTest2          = errors.New("error 2")
	errOriginal       = errors.New("original error")
	errFallback       = errors.New("fallback failed")
	testFieldName1    = "field1"
	testFieldName2    = "field2"
	testFailingField  = "failing_field"
	testFallbackValue = "fallback-value"
)

func newFluentWithConfig(cfg core.Configuration) (*zapfluent.Fluent, *zapcore.MapObjectEncoder) {
	enc := zapcore.NewMapObjectEncoder()
	return zapfluent.NewFluent(enc, cfg), enc
}

func TestFluent_Done_WithMultipleErrors_AggregatesErrors(t *testing.T) {
	g := NewWithT(t)
	fluent, _ := newFluentWithConfig(core.NewConfiguration())

	err := fluent.
		Add(stubs.NewFailingField(testFieldName1, errTest1)).
		Add(stubs.NewFailingField(testFieldName2, errTest2)).
		Done()

	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Is(err, errTest1)).To(BeTrue())
	g.Expect(errors.Is(err, errTest2)).To(BeTrue())
}

func TestFluent_ErrorHandling_EarlyFailing(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithMode(core.ErrorHandlingModeEarlyFailing),
			),
		),
	)
	fluent, _ := newFluentWithConfig(cfg)

	err := fluent.
		Add(stubs.NewFailingField(testFieldName1, errTest1)).
		Add(stubs.NewFailingField(testFieldName2, errTest2)).
		Done()

	g.Expect(errors.Is(err, errTest1)).To(BeTrue())
	g.Expect(errors.Is(err, errTest2)).To(BeFalse())
}

func TestFluent_WithFallback_ReplacesFailingFieldAndAggregatesError(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(core.FixedStringFallback(testFallbackValue)),
			),
		),
	)
	fluent, enc := newFluentWithConfig(cfg)

	err := fluent.
		Add(stubs.NewFailingField(testFailingField, errOriginal)).
		Done()

	g.Expect(errors.Is(err, errOriginal)).To(BeTrue())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, testFallbackValue))
}

func TestFluent_WithFailingFallback_LogsPredefinedErrorField(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingField(name, errFallback)
				}),
			),
		),
	)
	fluent, enc := newFluentWithConfig(cfg)

	err := fluent.
		Add(stubs.NewFailingField(testFailingField, errOriginal)).
		Done()

	g.Expect(errors.Is(err, errOriginal)).To(BeTrue())
	g.Expect(errors.Is(err, errFallback)).To(BeTrue())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "failed to encode fallback field"))
}

func TestFluent_WithFailingFallbackAndCustomMessage_LogsCustomMessage(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingField(name, errFallback)
				}),
				core.WithFallbackErrorMessage("custom message"),
			),
		),
	)
	fluent, enc := newFluentWithConfig(cfg)

	err := fluent.
		Add(stubs.NewFailingField(testFailingField, errOriginal)).
		Done()

	g.Expect(errors.Is(err, errOriginal)).To(BeTrue())
	g.Expect(errors.Is(err, errFallback)).To(BeTrue())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "custom message"))
}

func TestAsFluent_WithFluentEncoder(t *testing.T) {
	g := NewWithT(t)
	fluentEncoder := core.NewFluentEncoder(
		zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		core.NewConfiguration(),
	)

	fluent := zapfluent.AsFluent(fluentEncoder)

	g.Expect(fluent).ToNot(BeNil())
}

func TestAsFluent_WithOtherEncoder(t *testing.T) {
	g := NewWithT(t)
	fluent := zapfluent.AsFluent(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()))
	g.Expect(fluent).ToNot(BeNil())
}
