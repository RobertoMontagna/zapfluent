package zapfluent

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

func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return NewTypedField(
		objectTypeFns(isNonZero),
		name,
		value,
	)
}
