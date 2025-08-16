package zapfluent

import (
	"go.uber.org/zap/zapcore"
)

func Object[T zapcore.ObjectMarshaler](name string, value T, isNonZero func(T) bool) TypedField[T] {
	return NewTypedField(
		TypeFieldFunctions[T]{
			EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value T) error {
				return encoder.AddObject(name, value)
			},
			IsNonZero: isNonZero,
			FieldNoop: typedFieldNoop[T]{},
		},
		name,
		value,
	)
}
