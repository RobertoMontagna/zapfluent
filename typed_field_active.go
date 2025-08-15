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
	isPii     bool
}

func NewTypedField[T any](
	functions TypeFieldFunctions[T],
	value T,
) TypedField[T] {
	return TypedFieldActive[T]{
		functions: functions,
		value:     value,
	}
}

func (f TypedFieldActive[T]) Name(name string) zapcore.ObjectMarshalerFunc {
	f.name = name
	return f.encode
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
	return String(formatter(f.value))
}

func (f TypedFieldActive[T]) AsPII() TypedField[T] {
	f.isPii = true
	return f
}

func (f TypedFieldActive[T]) encode(encoder zapcore.ObjectEncoder) error {
	name := f.name
	if f.isPii {
		name += " <<pii>>"
	}
	return f.functions.EncodeFunc(encoder, name, f.value)
}
