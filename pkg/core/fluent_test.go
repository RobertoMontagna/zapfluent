package core_test

import (
	"errors"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
	"go.robertomontagna.dev/zapfluent/testutil"
	"go.robertomontagna.dev/zapfluent/testutil/stubs"

	. "github.com/onsi/gomega"
)

var (
	fErrTest1         = errors.New("error 1")
	fErrTest2         = errors.New("error 2")
	fErrOriginal      = errors.New("original error")
	fErrFallback      = errors.New("fallback failed")
	testFieldName1    = "field1"
	testFieldName2    = "field2"
	testFailingField  = "failing_field"
	testFallbackValue = "fallback-value"

	failingField1 = stubs.NewFailingFieldForTest(
		stubs.WithName(testFieldName1),
		stubs.WithError(fErrTest1),
	)
	failingField2 = stubs.NewFailingFieldForTest(
		stubs.WithName(testFieldName2),
		stubs.WithError(fErrTest2),
	)
	originalFailingField = stubs.NewFailingFieldForTest(
		stubs.WithName(testFailingField),
		stubs.WithError(fErrOriginal),
	)
)

func TestFluent_Done_WhenFieldsFail_ShouldAggregateErrors(t *testing.T) {
	g := NewWithT(t)

	fluent := core.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		core.NewConfiguration(),
	))

	err := fluent.
		Add(failingField1).
		Add(failingField2).
		Done()

	g.Expect(err).To(MatchError(fErrTest1))
	g.Expect(err).To(MatchError(fErrTest2))
}

func TestFluent_Done_WhenEarlyFailingIsEnabled_ShouldStopAfterFirstError(t *testing.T) {
	g := NewWithT(t)

	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithMode(core.ErrorHandlingModeEarlyFailing),
			),
		),
	)
	fluent := core.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		cfg,
	))

	err := fluent.
		Add(failingField1).
		Add(failingField2).
		Done()

	g.Expect(err).To(MatchError(fErrTest1))
	g.Expect(err).ToNot(MatchError(fErrTest2))
}

func TestFluent_Done_WhenFallbackIsConfigured_ShouldReplaceFieldAndReturnError(t *testing.T) {
	g := NewWithT(t)

	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(core.FixedStringFallback(testFallbackValue)),
			),
		),
	)
	enc := zapcore.NewMapObjectEncoder()
	fluent := core.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(fErrOriginal))
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, testFallbackValue))
}

func TestFluent_Done_WhenFallbackAlsoFails_ShouldLogPredefinedError(t *testing.T) {
	g := NewWithT(t)

	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingFieldForTest(
						stubs.WithName(name),
						stubs.WithError(fErrFallback),
					)
				}),
			),
		),
	)
	enc := zapcore.NewMapObjectEncoder()
	fluent := core.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(fErrOriginal))
	g.Expect(err).To(MatchError(fErrFallback))
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "failed to encode fallback field"))
}

func TestFluent_Done_WhenFailingFallbackHasCustomMessage_ShouldLogCustomMessage(t *testing.T) {
	g := NewWithT(t)

	cfg := core.NewConfiguration(
		core.WithErrorHandling(
			core.NewErrorHandlingConfiguration(
				core.WithFallbackFieldFactory(func(name string, err error) core.Field {
					return stubs.NewFailingFieldForTest(
						stubs.WithName(name),
						stubs.WithError(fErrFallback),
					)
				}),
				core.WithFallbackErrorMessage("custom message"),
			),
		),
	)
	enc := zapcore.NewMapObjectEncoder()
	fluent := core.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(fErrOriginal))
	g.Expect(err).To(MatchError(fErrFallback))
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "custom message"))
}

func TestAsFluent(t *testing.T) {
	testCases := []struct {
		name    string
		encoder zapcore.Encoder
	}{
		{
			name: "with fluent encoder",
			encoder: core.NewFluentEncoder(
				zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
				core.NewConfiguration(),
			),
		},
		{
			name:    "with other encoder",
			encoder: zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			fluent := core.AsFluent(tc.encoder)

			g.Expect(fluent).ToNot(BeNil())
		})
	}
}
