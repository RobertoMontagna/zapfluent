package zapfluent_test

import (
	"errors"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/pkg/core"
	"go.robertomontagna.dev/zapfluent/testutil"
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

func TestFluent_Done_WithMultipleErrors_AggregatesErrors(t *testing.T) {
	g := NewWithT(t)
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		core.NewConfiguration(),
	))

	err := fluent.
		Add(stubs.NewFailingFieldForTest(testFieldName1, errTest1)).
		Add(stubs.NewFailingFieldForTest(testFieldName2, errTest2)).
		Done()

	g.Expect(err).To(MatchError(errTest1))
	g.Expect(err).To(MatchError(errTest2))
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
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		cfg,
	))

	err := fluent.
		Add(stubs.NewFailingFieldForTest(testFieldName1, errTest1)).
		Add(stubs.NewFailingFieldForTest(testFieldName2, errTest2)).
		Done()

	g.Expect(err).To(MatchError(errTest1))
	g.Expect(err).ToNot(MatchError(errTest2))
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
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(stubs.NewFailingFieldForTest(testFailingField, errOriginal)).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, testFallbackValue))
}

func TestFluent_WithFailingFallback_LogsPredefinedErrorField(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingFieldForTest(name, errFallback)
				}),
			),
		),
	)
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(stubs.NewFailingFieldForTest(testFailingField, errOriginal)).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
	g.Expect(err).To(MatchError(errFallback))
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "failed to encode fallback field"))
}

func TestFluent_WithFailingFallbackAndCustomMessage_LogsCustomMessage(t *testing.T) {
	g := NewWithT(t)
	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingFieldForTest(name, errFallback)
				}),
				core.WithFallbackErrorMessage("custom message"),
			),
		),
	)
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(stubs.NewFailingFieldForTest(testFailingField, errOriginal)).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
	g.Expect(err).To(MatchError(errFallback))
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
