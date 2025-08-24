// Package stubs provides test stubs.
package stubs

import (
	"go.uber.org/zap/zapcore"
)

// FailingFieldForTest is a stub implementation of fluentfield.Field that is designed
// to always fail during encoding. This is useful for testing error-handling
// logic.
//
// This item is for testing purposes only and should not be used in production code.
type FailingFieldForTest struct {
	// Err is the error that will be returned when Encode is called.
	Err error
	// FieldName is the name of the field.
	FieldName string
}

// NewFailingFieldForTest creates a new FailingFieldForTest with the given name and error.
//
// This item is for testing purposes only and should not be used in production code.
func NewFailingFieldForTest(name string, err error) FailingFieldForTest {
	return FailingFieldForTest{
		FieldName: name,
		Err:       err,
	}
}

// Encode implements the fluentfield.Field interface and always returns the
// configured error.
func (f FailingFieldForTest) Encode(_ zapcore.ObjectEncoder) error {
	return f.Err
}

// Name returns the configured name of the field.
func (f FailingFieldForTest) Name() string {
	if f.FieldName == "" {
		return "error"
	}
	return f.FieldName
}
