//go:build test

// Package testutil provides helpers for testing purposes.
package testutil

import "go.uber.org/zap/zapcore"

// FailingField is a mock implementation of fluentfield.Field that is designed
// to always fail during encoding. This is useful for testing error-handling
// logic.
type FailingField struct {
	// Err is the error that will be returned when Encode is called.
	Err error
	// NameValue is the name of the field.
	NameValue string
}

// Encode implements the fluentfield.Field interface and always returns the
// configured error.
func (f FailingField) Encode(_ zapcore.ObjectEncoder) error {
	return f.Err
}

// Name returns the configured name of the field.
func (f FailingField) Name() string {
	if f.NameValue == "" {
		return "error"
	}
	return f.NameValue
}
