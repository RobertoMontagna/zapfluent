package zapfluent

import "go.uber.org/zap/zapcore"

type EncodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

type TypeFieldFunctions[T any] struct {
	EncodeFunc EncodeFunc[T]
	IsNonZero  func(T) bool
	FieldNoop  TypedField[T]
}

type TypedFieldActive[T any] struct {
	functions TypeFieldFunctions[T]
	value     T
	name      string
}

func NewTypedField[T any](
	functions TypeFieldFunctions[T],
	name string,
	value T,
) TypedField[T] {
	return TypedFieldActive[T]{
		functions: functions,
		name:      name,
		value:     value,
	}
}

func (f TypedFieldActive[T]) Filter(condition func(T) bool) TypedField[T] {
	if condition(f.value) {
		return f
	}
	return f.functions.FieldNoop
}

func (f TypedFieldActive[T]) NonZero() TypedField[T] {
	return f.Filter(f.functions.IsNonZero)
}

func (f TypedFieldActive[T]) Format(formatter func(T) string) TypedField[string] {
	return String(f.name, formatter(f.value))
}

func (f TypedFieldActive[T]) Encode(encoder zapcore.ObjectEncoder) error {
	return f.functions.EncodeFunc(encoder, f.name, f.value)
}
