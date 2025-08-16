package zapfluent

import "go.uber.org/zap/zapcore"

type comparableStruct interface {
	zapcore.ObjectMarshaler
	comparable
}

func ComparableStruct[T comparableStruct](name string, value T) TypedField[T] {
	return NewTypedField(
		TypeFieldFunctions[T]{
			EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value T) error {
				return encoder.AddObject(name, value)
			},
			IsNonZero: func(v T) bool {
				var x T
				return v != x
			},
			FieldNoop: typedFieldNoop[T]{},
		},
		name,
		value,
	)
}
