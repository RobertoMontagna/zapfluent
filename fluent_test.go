package zapfluent_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/internal/testutil"

	. "github.com/onsi/gomega"
)

const (
	testError1        = "error 1"
	testError2        = "error 2"
	testFieldName1    = "field1"
	testFieldName2    = "field2"
	testOriginalError = "original error"
	testFallbackValue = "fallback-value"
	testFailingField  = "failing_field"
	testFallbackError = "fallback failed"
)

func newFluentWithConfig(cfg zapfluent.Configuration) (*zapfluent.Fluent, *zapcore.MapObjectEncoder) {
	enc := zapcore.NewMapObjectEncoder()
	fluent := zapfluent.NewFluent(enc, cfg)
	return fluent, enc
}

func TestFluent(t *testing.T) {
	err1 := errors.New(testError1)
	err2 := errors.New(testError2)
	originalErr := errors.New(testOriginalError)
	fallbackErr := errors.New(testFallbackError)

	scenarios := []struct {
		name       string
		cfg        zapfluent.Configuration
		fields     []zapfluent.Field
		assertions func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder)
	}{
		{
			name:   "Done_WithMultipleErrors_AggregatesErrors",
			cfg:    zapfluent.NewConfiguration(),
			fields: []zapfluent.Field{testutil.FailingField{Err: err1}, testutil.FailingField{Err: err2}},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(testError1))
				g.Expect(err.Error()).To(ContainSubstring(testError2))
			},
		},
		{
			name: "ErrorHandling_EarlyFailing",
			cfg: zapfluent.NewConfiguration(
				zapfluent.WithErrorHandling(
					zapfluent.NewErrorHandlingConfiguration(
						zapfluent.WithMode(zapfluent.ErrorHandlingModeEarlyFailing),
					),
				),
			),
			fields: []zapfluent.Field{
				testutil.FailingField{Err: err1, NameValue: testFieldName1},
				testutil.FailingField{Err: err2, NameValue: testFieldName2},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(err1)) // Should only return the first error
			},
		},
		{
			name: "WithFallback_replaces_the_failing_field_and_aggregates_the_error",
			cfg: zapfluent.NewConfiguration(
				zapfluent.WithErrorHandling(
					zapfluent.NewErrorHandlingConfiguration(
						zapfluent.WithFallbackFieldFactory(zapfluent.FixedStringFallback(testFallbackValue)),
					),
				),
			),
			fields: []zapfluent.Field{
				testutil.FailingField{Err: originalErr, NameValue: testFailingField},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(originalErr))
				g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, testFallbackValue))
			},
		},
		{
			name: "WithFailingFallback_logs_a_predefined_error_field",
			cfg: zapfluent.NewConfiguration(
				zapfluent.WithErrorHandling(
					zapfluent.NewErrorHandlingConfiguration(
						zapfluent.WithFallbackFieldFactory(func(name string, err error) zapfluent.Field {
							return testutil.FailingField{NameValue: name, Err: fallbackErr}
						}),
					),
				),
			),
			fields: []zapfluent.Field{
				testutil.FailingField{Err: originalErr, NameValue: testFailingField},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(originalErr))
				g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "failed to encode fallback field"))
			},
		},
		{
			name: "WithFailingFallback_and_custom_message_logs_the_custom_message",
			cfg: zapfluent.NewConfiguration(
				zapfluent.WithErrorHandling(
					zapfluent.NewErrorHandlingConfiguration(
						zapfluent.WithFallbackFieldFactory(func(name string, err error) zapfluent.Field {
							return testutil.FailingField{NameValue: name, Err: fallbackErr}
						}),
						zapfluent.WithFallbackErrorMessage("custom message"),
					),
				),
			),
			fields: []zapfluent.Field{
				testutil.FailingField{Err: originalErr, NameValue: testFailingField},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(originalErr))
				g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, "custom message"))
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			g := NewWithT(t)

			fluent, enc := newFluentWithConfig(s.cfg)
			for _, f := range s.fields {
				fluent.Add(f)
			}

			err := fluent.Done()

			s.assertions(g, err, enc)
		})
	}
}
