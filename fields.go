package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// Field is the interface that all concrete field types must implement. It
// represents a single key-value pair to be encoded.
type Field = core.Field

// TypedField is a generic interface that extends the basic Field interface with
// methods for filtering and transformation. This allows for creating expressive,
// chainable APIs for constructing fields.
type TypedField[T any] = core.TypedField[T]

// String returns a new field with a string value.
func String(name string, value string) TypedField[string] {
	return core.String(name, value)
}

// Int returns a new field with an int value.
func Int(name string, value int) TypedField[int] {
	return core.Int(name, value)
}

// Int8 returns a new field with an int8 value.
func Int8(name string, value int8) TypedField[int8] {
	return core.Int8(name, value)
}

// Object returns a new field with a value that implements zapcore.ObjectMarshaler.
//
// It requires an `isNonZero` function to determine if the object should be
// omitted when the `NonZero` method is called.
func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return core.Object(name, value, isNonZero)
}

// ComparableObject returns a new field with a value that implements both
// zapcore.ObjectMarshaler and the comparable constraint.
//
// The `isNonZero` function for this field performs a simple comparison to the
// zero value of the type (e.g., `v != *new(T)`).
func ComparableObject[T core.Comparable](name string, value T) TypedField[T] {
	return core.ComparableObject(name, value)
}
