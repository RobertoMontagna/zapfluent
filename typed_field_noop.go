package zapfluent

import "go.uber.org/zap/zapcore"

type typedFieldNoop[T any] struct {
}

func (f typedFieldNoop[T]) Encode(_ zapcore.ObjectEncoder) error {
	return nil
}

func (f typedFieldNoop[T]) Filter(_ func(T) bool) TypedField[T] {
	return f
}

func (f typedFieldNoop[T]) NonZero() TypedField[T] {
	return f
}

func (f typedFieldNoop[T]) Format(_ func(T) string) TypedField[string] {
	return typedFieldNoop[string]{}
}
