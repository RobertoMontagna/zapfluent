package zapfluent

import "go.uber.org/zap/zapcore"

type comparableStruct interface {
	zapcore.ObjectMarshaler
	comparable
}

func comparableStructFunctions[T comparableStruct]() TypeFieldFunctions[T] {
	return TypeFieldFunctions[T]{
		EncodeFunc: comparableStructEncodeFunc[T](),
		IsNonZero:  comparableStructIsNonZero[T](),
		FieldNoop:  typedFieldNoop[T]{},
	}
}

func comparableStructEncodeFunc[T comparableStruct]() func(encoder zapcore.ObjectEncoder, name string, value T) error {
	return func(encoder zapcore.ObjectEncoder, name string, value T) error {
		return encoder.AddObject(name, value)
	}
}

func comparableStructIsNonZero[T comparableStruct]() func(value T) bool {
	return func(value T) bool {
		var x T
		return value != x
	}
}

func ComparableStruct[T comparableStruct](name string, value T) TypedField[T] {
	return NewTypedField(
		comparableStructFunctions[T](),
		name,
		value,
	)
}
