package fluentfield

import (
	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
	"go.uber.org/zap/zapcore"
)

type EncodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

type TypeFieldFunctions[T any] struct {
	EncodeFunc EncodeFunc[T]
	IsNonZero  func(T) bool
}

type LazyTypedField[T any] struct {
	functions TypeFieldFunctions[T]
	optional  lazyoptional.LazyOptional[T]
	name      string
}

func NewTypedField[T any](
	functions TypeFieldFunctions[T],
	name string,
	value T,
) TypedField[T] {
	return &LazyTypedField[T]{
		functions: functions,
		name:      name,
		optional:  lazyoptional.Some(value),
	}
}

func (f *LazyTypedField[T]) Name() string {
	return f.name
}

func (f *LazyTypedField[T]) Encode(encoder zapcore.ObjectEncoder) error {
	val, ok := f.optional.Get()
	if !ok {
		return nil
	}
	return f.functions.EncodeFunc(encoder, f.name, val)
}

func (f *LazyTypedField[T]) Filter(condition func(T) bool) TypedField[T] {
	return &LazyTypedField[T]{
		functions: f.functions,
		name:      f.name,
		optional:  f.optional.Filter(condition),
	}
}

func (f *LazyTypedField[T]) NonZero() TypedField[T] {
	return f.Filter(f.functions.IsNonZero)
}

func (f *LazyTypedField[T]) Format(formatter func(T) string) TypedField[string] {
	return &LazyTypedField[string]{
		name:      f.name,
		functions: stringTypeFns(),
		optional:  lazyoptional.Map(f.optional, formatter),
	}
}
