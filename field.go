package zapfluent

import "go.uber.org/zap/zapcore"

type Field interface {
	Encode(zapcore.ObjectEncoder) error
}

type TypedField[T any] interface {
	Field
	Filter(condition func(T) bool) TypedField[T]
	NonZero() TypedField[T]
	Format(formatter func(T) string) TypedField[string]
}
