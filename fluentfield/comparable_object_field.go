package fluentfield

import "go.uber.org/zap/zapcore"

type comparableObject interface {
	zapcore.ObjectMarshaler
	comparable
}

// ComparableObject returns a new field with a value that implements both
// zapcore.ObjectMarshaler and the comparable constraint.
//
// The `IsNonZero` function for this field performs a simple comparison to the
// zero value of the type (e.g., `v != *new(T)`).
func ComparableObject[T comparableObject](name string, value T) TypedField[T] {
	return Object(name, value, func(v T) bool {
		var x T
		return v != x
	})
}
