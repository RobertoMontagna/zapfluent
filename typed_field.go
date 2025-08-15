package zapfluent

import "go.uber.org/zap/zapcore"

type TypedField[T any] interface {
	Name(name string) zapcore.ObjectMarshalerFunc
	Filter(condition func(T) bool) TypedField[T]
	NonZero() TypedField[T]
	Format(formatter func(T) string) TypedField[string]
	AsPII() TypedField[T]
}
