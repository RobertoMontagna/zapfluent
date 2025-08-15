package zapfluent

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// TODO add all the remaining types
type primitiveType interface {
	string | ~int | ~int8
}

type PrimitiveFunctionsOption[T primitiveType] func(*TypeFieldFunctions[T])

func primitiveFunctions[T primitiveType](
	options ...PrimitiveFunctionsOption[T],
) TypeFieldFunctions[T] {
	config := &TypeFieldFunctions[T]{
		EncodeFunc: primitiveEncodeFunc[T](),
		IsNonZero:  primitiveIsNonZero[T](),
		FieldNoop:  typedFieldNoop[T]{},
	}
	for _, option := range options {
		option(config)
	}
	return *config
}

func primitiveEncodeFunc[T primitiveType]() func(encoder zapcore.ObjectEncoder, name string, value T) error {
	return func(encoder zapcore.ObjectEncoder, name string, value T) error {
		switch v := any(value).(type) {
		case string:
			encoder.AddString(name, v)
		case int:
			encoder.AddInt(name, v)
		case int8:
			encoder.AddInt8(name, v)
		default:
			panic(fmt.Errorf("unsupported primitive type: %T", v))
		}
		return nil
	}
}

func primitiveIsNonZero[T primitiveType]() func(value T) bool {
	return func(value T) bool {
		var x T
		return value != x
	}
}

func Primitive[T primitiveType](
	options ...PrimitiveFunctionsOption[T],
) func(value T) TypedField[T] {
	return func(value T) TypedField[T] {
		return NewTypedField(
			primitiveFunctions[T](options...),
			value,
		)
	}
}
