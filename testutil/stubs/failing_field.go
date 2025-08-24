// Package stubs provides test stubs.
package stubs

import (
	"go.uber.org/zap/zapcore"
)

// FailingField is a stub implementation of fluentfield.Field that is designed
// to always fail during encoding. This is useful for testing error-handling
// logic.
type FailingField struct {
	// Err is the error that will be returned when Encode is called.
	Err error
	// FieldName is the name of the field.
	FieldName string
}

// NewFailingField creates a new FailingField with the given name and error.
func NewFailingField(name string, err error) FailingField {
	return FailingField{
		FieldName: name,
		Err:       err,
	}
}

// Encode implements the fluentfield.Field interface and always returns the
// configured error.
func (f FailingField) Encode(_ zapcore.ObjectEncoder) error {
	return f.Err
}

// Name returns the configured name of the field.
func (f FailingField) Name() string {
	if f.FieldName == "" {
		return "error"
	}
	return f.FieldName
}
