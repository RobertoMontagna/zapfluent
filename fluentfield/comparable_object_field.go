package fluentfield

import "go.uber.org/zap/zapcore"

type comparableObject interface {
	zapcore.ObjectMarshaler
	comparable
}

func ComparableObject[T comparableObject](name string, value T) TypedField[T] {
	return Object(name, value, func(v T) bool {
		var x T
		return v != x
	})
}
