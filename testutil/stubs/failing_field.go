// Package stubs provides test stubs.
package stubs

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// FailingFieldForTest is a stub implementation of fluentfield.Field that is designed
// to always fail during encoding. This is useful for testing error-handling
// logic.
//
// This item is for testing purposes only and should not be used in production code.
type FailingFieldForTest struct {
	err  error
	name string
}

// FailingFieldForTestOption is an option for NewFailingFieldForTest.
//
// This item is for testing purposes only and should not be used in production code.
type FailingFieldForTestOption func(*FailingFieldForTest)

// WithName sets the name of the field.
// It panics if the name is empty.
//
// This item is for testing purposes only and should not be used in production code.
func WithName(name string) FailingFieldForTestOption {
	if name == "" {
		panic("name cannot be empty")
	}

	return func(f *FailingFieldForTest) {
		f.name = name
	}
}

// WithError sets the error that will be returned when Encode is called.
// It panics if the error is nil.
//
// This item is for testing purposes only and should not be used in production code.
func WithError(err error) FailingFieldForTestOption {
	if err == nil {
		panic("error cannot be nil")
	}

	return func(f *FailingFieldForTest) {
		f.err = err
	}
}

// NewFailingFieldForTest creates a new FailingFieldForTest with the given options.
//
// This item is for testing purposes only and should not be used in production code.
func NewFailingFieldForTest(opts ...FailingFieldForTestOption) FailingFieldForTest {
	sut := &FailingFieldForTest{
		name: "error",
		err:  fmt.Errorf("unspecified error"),
	}

	for _, opt := range opts {
		opt(sut)
	}

	return *sut
}

// Encode implements the fluentfield.Field interface and always returns the
// configured error.
func (f FailingFieldForTest) Encode(_ zapcore.ObjectEncoder) error {
	return f.err
}

// Name returns the configured name of the field.
func (f FailingFieldForTest) Name() string {
	return f.name
}
