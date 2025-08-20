package fluentfield

import (
	"go.uber.org/zap/zapcore"
)

func objectTypeFns[T zapcore.ObjectMarshaler](isNonZero func(T) bool) TypeFieldFunctions[T] {
	return TypeFieldFunctions[T]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value T) error {
			return encoder.AddObject(name, value)
		},
		IsNonZero: isNonZero,
	}
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
