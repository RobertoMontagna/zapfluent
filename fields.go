package zapfluent

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"
)

type Fluent = core.Fluent

// AsFluent returns a Fluent wrapper for the provided zapcore.ObjectEncoder,
// enabling the fluent field-construction helpers in this package.
func AsFluent(encoder zapcore.ObjectEncoder) *Fluent {
	return core.AsFluent(encoder)
}

// Field is the interface that all concrete field types must implement. It
// represents a single key-value pair to be encoded.
type Field = core.Field

// TypedField is a generic interface that extends the basic Field interface with
// methods for filtering and transformation. This allows for creating expressive,
// chainable APIs for constructing fields.
type TypedField[T any] = core.TypedField[T]

// TypedPointerField is a generic interface that represents a field with a
// pointer value. It implements the base Field interface and provides a `NonNil`
// method to safely convert it to a TypedField for chaining.
type TypedPointerField[T any] = core.TypedPointerField[T]

// String returns a new field with a string value.
func String(name string, value string) TypedField[string] {
	return core.String(name, value)
}

// StringPtr returns a new field with a *string value.
func StringPtr(name string, value *string) TypedPointerField[string] {
	return core.StringPtr(name, value)
}

// Int returns a new field with an int value.
func Int(name string, value int) TypedField[int] {
	return core.Int(name, value)
}

// IntPtr returns a new field with an *int value.
func IntPtr(name string, value *int) TypedPointerField[int] {
	return core.IntPtr(name, value)
}

// Int8 returns a new field with an int8 value.
func Int8(name string, value int8) TypedField[int8] {
	return core.Int8(name, value)
}

// Int8Ptr returns a new field with an *int8 value.
func Int8Ptr(name string, value *int8) TypedPointerField[int8] {
	return core.Int8Ptr(name, value)
}

// Object returns a new field with a value that implements zapcore.ObjectMarshaler.
//
// It requires an `isNonZero` function to determine if the object should be
// omitted when the `NonZero` method is called.
func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return core.Object(name, value, isNonZero)
}

// ObjectPtr returns a new field with a value that is a pointer to a
// zapcore.ObjectMarshaler.
func ObjectPtr[T zapcore.ObjectMarshaler](
	name string,
	value *T,
	isNonZero func(T) bool,
) TypedPointerField[T] {
	return core.ObjectPtr(name, value, isNonZero)
}

// ComparableObject returns a new field with a value that implements both
// zapcore.ObjectMarshaler and the comparable constraint.
//
// The `isNonZero` function for this field performs a simple comparison to the
// zero value of the type (e.g., `v != *new(T)`).
func ComparableObject[T core.Comparable](name string, value T) TypedField[T] {
	return core.ComparableObject(name, value)
}

// ComparableObjectPtr returns a new field with a value that is a pointer to a
// type that implements both zapcore.ObjectMarshaler and the comparable constraint.
func ComparableObjectPtr[T core.Comparable](name string, value *T) TypedPointerField[T] {
	return core.ComparableObjectPtr(name, value)
}

// Bool returns a new field with a bool value.
func Bool(name string, value bool) TypedField[bool] {
	return core.Bool(name, value)
}

// BoolPtr returns a new field with a *bool value.
func BoolPtr(name string, value *bool) TypedPointerField[bool] {
	return core.BoolPtr(name, value)
}
