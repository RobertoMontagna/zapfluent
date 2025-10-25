package core

import "go.uber.org/zap/zapcore"

var (
	// boolTypeFns holds the cached typeFieldFunctions for bool fields.
	boolTypeFns = typeFieldFunctions[bool]{
		encodeFunc: func(encoder zapcore.ObjectEncoder, name string, value bool) error {
			encoder.AddBool(name, value)
			return nil
		},
		isNonZero: func(b bool) bool {
			return b
		},
	}
	boolTypePointerFns = primitiveTypePointerFns(boolTypeFns, Bool)

	// intTypeFns holds the cached typeFieldFunctions for int fields.
	intTypeFns = typeFieldFunctions[int]{
		encodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int) error {
			encoder.AddInt(name, value)
			return nil
		},
		isNonZero: func(i int) bool {
			return i != 0
		},
	}
	intTypePointerFns = primitiveTypePointerFns(intTypeFns, Int)

	// int8TypeFns holds the cached typeFieldFunctions for int8 fields.
	int8TypeFns = typeFieldFunctions[int8]{
		encodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int8) error {
			encoder.AddInt8(name, value)
			return nil
		},
		isNonZero: func(i int8) bool {
			return i != 0
		},
	}
	int8TypePointerFns = primitiveTypePointerFns(int8TypeFns, Int8)

	// stringTypeFns holds the cached typeFieldFunctions for string fields.
	stringTypeFns = typeFieldFunctions[string]{
		encodeFunc: func(encoder zapcore.ObjectEncoder, name string, value string) error {
			encoder.AddString(name, value)
			return nil
		},
		isNonZero: func(s string) bool {
			return s != ""
		},
	}
	stringTypePointerFns = primitiveTypePointerFns(stringTypeFns, String)
)

// Bool returns a new field with a bool value.
func Bool(name string, value bool) TypedField[bool] {
	return newTypedField(
		boolTypeFns,
		name,
		value,
	)
}

// BoolPtr returns a new field with a *bool value.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func BoolPtr(name string, value *bool) TypedPointerField[bool] {
	return newPointerField(
		boolTypePointerFns,
		name,
		value,
	)
}

// Int returns a new field with an int value.
func Int(name string, value int) TypedField[int] {
	return newTypedField(
		intTypeFns,
		name,
		value,
	)
}

// IntPtr returns a new field with an *int value.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func IntPtr(name string, value *int) TypedPointerField[int] {
	return newPointerField(
		intTypePointerFns,
		name,
		value,
	)
}

// Int8 returns a new field with an int8 value.
func Int8(name string, value int8) TypedField[int8] {
	return newTypedField(
		int8TypeFns,
		name,
		value,
	)
}

// Int8Ptr returns a new field with an *int8 value.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func Int8Ptr(name string, value *int8) TypedPointerField[int8] {
	return newPointerField(
		int8TypePointerFns,
		name,
		value,
	)
}

// String returns a new field with a string value.
func String(name string, value string) TypedField[string] {
	return newTypedField(
		stringTypeFns,
		name,
		value,
	)
}

// StringPtr returns a new field with a *string value.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func StringPtr(name string, value *string) TypedPointerField[string] {
	return newPointerField(
		stringTypePointerFns,
		name,
		value,
	)
}

// Object returns a new field with a value that implements zapcore.ObjectMarshaler.
//
// It requires an `isNonZero` function to determine if the object should be
// omitted when the `NonZero` method is called.
func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return newTypedField(
		objectTypeFns(isNonZero),
		name,
		value,
	)
}

// ObjectPtr returns a new field with a value that is a pointer to a zapcore.ObjectMarshaler.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func ObjectPtr[T zapcore.ObjectMarshaler](
	name string,
	value *T,
	isNonZero func(T) bool,
) TypedPointerField[T] {
	return newPointerField(
		objectTypePointerFns(isNonZero),
		name,
		value,
	)
}

type Comparable interface {
	zapcore.ObjectMarshaler
	comparable
}

// ComparableObject returns a new field with a value that implements both
// zapcore.ObjectMarshaler and the comparable constraint.
//
// The `isNonZero` function for this field performs a simple comparison to the
// zero value of the type (e.g., `v != *new(T)`).
func ComparableObject[T Comparable](name string, value T) TypedField[T] {
	var zero T
	return Object(name, value, func(v T) bool {
		return v != zero
	})
}

// ComparableObjectPtr returns a new field with a value that is a pointer to a
// type that implements both zapcore.ObjectMarshaler and the comparable constraint.
// When value is nil, it encodes as the string sentinel (e.g., "<nil>") to make nil explicit.
func ComparableObjectPtr[T Comparable](name string, value *T) TypedPointerField[T] {
	var zero T
	return ObjectPtr(name, value, func(v T) bool {
		return v != zero
	})
}

func primitiveTypePointerFns[T any](
	base typeFieldFunctions[T],
	toField func(string, T) TypedField[T],
) typePointerFieldFunctions[T] {
	return typePointerFieldFunctions[T]{
		typeFieldFunctions: base,
		toField:            toField,
	}
}

func objectTypeFns[T zapcore.ObjectMarshaler](isNonZero func(T) bool) typeFieldFunctions[T] {
	return typeFieldFunctions[T]{
		encodeFunc: func(encoder zapcore.ObjectEncoder, name string, value T) error {
			return encoder.AddObject(name, value)
		},
		isNonZero: isNonZero,
	}
}

func objectTypePointerFns[T zapcore.ObjectMarshaler](isNonZero func(T) bool) typePointerFieldFunctions[T] {
	return typePointerFieldFunctions[T]{
		typeFieldFunctions: objectTypeFns(isNonZero),
		toField: func(name string, value T) TypedField[T] {
			return Object(name, value, isNonZero)
		},
	}
}
