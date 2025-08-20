package zapfluent_test

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/config"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/util/testing_util"
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

func newFluentWithConfig(cfg config.Configuration) (*zapfluent.Fluent, *zapcore.MapObjectEncoder) {
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
		cfg        config.Configuration
		fields     []fluentfield.Field
		assertions func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder)
	}{
		{
			name:   "Done_WithMultipleErrors_AggregatesErrors",
			cfg:    config.NewConfiguration(),
			fields: []fluentfield.Field{testing_util.FailingField{Err: err1}, testing_util.FailingField{Err: err2}},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(testError1))
				g.Expect(err.Error()).To(ContainSubstring(testError2))
			},
		},
		{
			name: "ErrorHandling_EarlyFailing",
			cfg: config.NewConfiguration(
				config.WithErrorHandling(
					config.NewErrorHandlingConfiguration(
						config.WithMode(config.ErrorHandlingModeEarlyFailing),
					),
				),
			),
			fields: []fluentfield.Field{
				testing_util.FailingField{Err: err1, NameValue: testFieldName1},
				testing_util.FailingField{Err: err2, NameValue: testFieldName2},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(err1)) // Should only return the first error
			},
		},
		{
			name: "WithFallback_replaces_the_failing_field_and_aggregates_the_error",
			cfg: config.NewConfiguration(
				config.WithErrorHandling(
					config.NewErrorHandlingConfiguration(
						config.WithFallbackFieldFactory(config.FixedStringFallback(testFallbackValue)),
					),
				),
			),
			fields: []fluentfield.Field{
				testing_util.FailingField{Err: originalErr, NameValue: testFailingField},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(MatchError(originalErr))
				g.Expect(enc.Fields).To(HaveKeyWithValue(testFailingField, testFallbackValue))
			},
		},
		{
			name: "WithFallback_aggregates_errors_from_the_fallback_field_itself",
			cfg: config.NewConfiguration(
				config.WithErrorHandling(
					config.NewErrorHandlingConfiguration(
						config.WithFallbackFieldFactory(func(name string, err error) fluentfield.Field {
							return testing_util.FailingField{NameValue: name, Err: fallbackErr}
						}),
					),
				),
			),
			fields: []fluentfield.Field{
				testing_util.FailingField{Err: originalErr, NameValue: testFailingField},
			},
			assertions: func(g *GomegaWithT, err error, enc *zapcore.MapObjectEncoder) {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(originalErr.Error()))
				g.Expect(err.Error()).To(ContainSubstring(fallbackErr.Error()))
				g.Expect(enc.Fields).To(BeEmpty())
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
