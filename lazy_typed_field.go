package zapfluent

import "go.uber.org/zap/zapcore"

type EncodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

type TypeFieldFunctions[T any] struct {
	EncodeFunc EncodeFunc[T]
	IsNonZero  func(T) bool
	FieldNoop  TypedField[T]
}

type LazyValue[T any] func() (T, bool)

type LazyTypedField[T any] struct {
	functions TypeFieldFunctions[T]
	value     LazyValue[T]
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
		value: func() (T, bool) {
			return value, true
		},
	}
}

func (f *LazyTypedField[T]) Encode(encoder zapcore.ObjectEncoder) error {
	val, ok := f.value()
	if !ok {
		return nil
	}
	return f.functions.EncodeFunc(encoder, f.name, val)
}

func (f *LazyTypedField[T]) Filter(condition func(T) bool) TypedField[T] {
	return &LazyTypedField[T]{
		functions: f.functions,
		name:      f.name,
		value: func() (T, bool) {
			val, ok := f.value()
			if !ok {
				return val, false
			}
			if condition(val) {
				return val, true
			}
			var zero T
			return zero, false
		},
	}
}

func (f *LazyTypedField[T]) NonZero() TypedField[T] {
	return f.Filter(f.functions.IsNonZero)
}

func (f *LazyTypedField[T]) Format(formatter func(T) string) TypedField[string] {
	return &LazyTypedField[string]{
		name:      f.name,
		functions: stringTypeFns(),
		value: func() (string, bool) {
			val, ok := f.value()
			if !ok {
				var zero string
				return zero, false
			}
			return formatter(val), true
		},
	}
}
