package fluentfield

import "go.uber.org/zap/zapcore"

// String returns a new field with a string value.
func String(name string, value string) TypedField[string] {
	return NewTypedField(
		stringTypeFns(),
		name,
		value,
	)
}

// Int returns a new field with an int value.
func Int(name string, value int) TypedField[int] {
	return NewTypedField(
		intTypeFns(),
		name,
		value,
	)
}

// Int8 returns a new field with an int8 value.
func Int8(name string, value int8) TypedField[int8] {
	return NewTypedField(
		int8TypeFns(),
		name,
		value,
	)
}

// Object returns a new field with a value that implements zapcore.ObjectMarshaler.
//
// It requires an `isNonZero` function to determine if the object should be
// omitted when the `NonZero` method is called.
func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return NewTypedField(
		objectTypeFns(isNonZero),
		name,
		value,
	)
}

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

// unexported helpers that were in the original files
func stringTypeFns() TypeFieldFunctions[string] {
	return TypeFieldFunctions[string]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value string) error {
			encoder.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool {
			return s != ""
		},
	}
}

func intTypeFns() TypeFieldFunctions[int] {
	return TypeFieldFunctions[int]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int) error {
			encoder.AddInt(name, value)
			return nil
		},
		IsNonZero: func(i int) bool {
			return i != 0
		},
	}
}

func int8TypeFns() TypeFieldFunctions[int8] {
	return TypeFieldFunctions[int8]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int8) error {
			encoder.AddInt8(name, value)
			return nil
		},
		IsNonZero: func(i int8) bool {
			return i != 0
		},
	}
}

func objectTypeFns[T zapcore.ObjectMarshaler](isNonZero func(T) bool) TypeFieldFunctions[T] {
	return TypeFieldFunctions[T]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value T) error {
			return encoder.AddObject(name, value)
		},
		IsNonZero: isNonZero,
	}
}
