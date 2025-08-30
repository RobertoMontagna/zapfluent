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

	failingField1 = stubs.NewFailingFieldForTest(
		stubs.WithName(testFieldName1),
		stubs.WithError(errTest1),
	)
	failingField2 = stubs.NewFailingFieldForTest(
		stubs.WithName(testFieldName2),
		stubs.WithError(errTest2),
	)
	originalFailingField = stubs.NewFailingFieldForTest(
		stubs.WithName(testFailingField),
		stubs.WithError(errOriginal),
	)
)

func TestFluent_Done_WhenFieldsFail_ShouldAggregateErrors(t *testing.T) {
	g := NewWithT(t)

	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		core.NewConfiguration(),
	))

	err := fluent.
		Add(failingField1).
		Add(failingField2).
		Done()

	g.Expect(err).To(MatchError(errTest1))
	g.Expect(err).To(MatchError(errTest2))
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
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(zapcore.NewMapObjectEncoder()),
		cfg,
	))

	err := fluent.
		Add(failingField1).
		Add(failingField2).
		Done()

	g.Expect(err).To(MatchError(errTest1))
	g.Expect(err).ToNot(MatchError(errTest2))
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
	fluent := zapfluent.AsFluent(core.NewFluentEncoder(
		testutil.NewDoNotEncodeEncoderForTest(enc),
		cfg,
	))

	err := fluent.
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
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
						stubs.WithError(errFallback),
					)
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
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
	g.Expect(err).To(MatchError(errFallback))
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
						stubs.WithError(errFallback),
					)
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
		Add(originalFailingField).
		Done()

	g.Expect(err).To(MatchError(errOriginal))
	g.Expect(err).To(MatchError(errFallback))
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

			fluent := zapfluent.AsFluent(tc.encoder)

			g.Expect(fluent).ToNot(BeNil())
		})
	}
}

func TestFluent_Add_ForDifferentFieldTypes_ShouldEncodeCorrectly(t *testing.T) {
	strVal := "value"
	intVal := 42
	int8Val := int8(8)
	boolTrue := true

	testCases := []struct {
		name          string
		field         zapfluent.Field
		expectedKey   string
		expectedValue any
	}{
		{
			name:          "with string",
			field:         zapfluent.String("my_string", "value"),
			expectedKey:   "my_string",
			expectedValue: "value",
		},
		{
			name:          "with string pointer (non-nil)",
			field:         zapfluent.StringPtr("my_string_ptr", &strVal),
			expectedKey:   "my_string_ptr",
			expectedValue: "value",
		},
		{
			name:          "with string pointer (nil)",
			field:         zapfluent.StringPtr("my_string_ptr", nil),
			expectedKey:   "my_string_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with int",
			field:         zapfluent.Int("my_int", 123),
			expectedKey:   "my_int",
			expectedValue: 123,
		},
		{
			name:          "with int pointer (non-nil)",
			field:         zapfluent.IntPtr("my_int_ptr", &intVal),
			expectedKey:   "my_int_ptr",
			expectedValue: 42,
		},
		{
			name:          "with int pointer (nil)",
			field:         zapfluent.IntPtr("my_int_ptr", nil),
			expectedKey:   "my_int_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with int8",
			field:         zapfluent.Int8("my_int8", 12),
			expectedKey:   "my_int8",
			expectedValue: int8(12),
		},
		{
			name:          "with int8 pointer (non-nil)",
			field:         zapfluent.Int8Ptr("my_int8_ptr", &int8Val),
			expectedKey:   "my_int8_ptr",
			expectedValue: int8(8),
		},
		{
			name:          "with int8 pointer (nil)",
			field:         zapfluent.Int8Ptr("my_int8_ptr", nil),
			expectedKey:   "my_int8_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with bool (true)",
			field:         zapfluent.Bool("my_bool", true),
			expectedKey:   "my_bool",
			expectedValue: true,
		},
		{
			name:          "with bool pointer (non-nil)",
			field:         zapfluent.BoolPtr("my_bool_ptr", &boolTrue),
			expectedKey:   "my_bool_ptr",
			expectedValue: true,
		},
		{
			name:          "with bool pointer (nil)",
			field:         zapfluent.BoolPtr("my_bool_ptr", nil),
			expectedKey:   "my_bool_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with bool (false)",
			field:         zapfluent.Bool("my_bool", false),
			expectedKey:   "my_bool",
			expectedValue: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			fluent := zapfluent.AsFluent(
				core.NewFluentEncoder(
					testutil.NewDoNotEncodeEncoderForTest(enc),
					core.NewConfiguration(),
				),
			)

			err := fluent.Add(tc.field).Done()

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
		})
	}
}
