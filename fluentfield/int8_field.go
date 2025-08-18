package fluentfield

import "go.uber.org/zap/zapcore"

func int8TypeFns() TypeFieldFunctions[int8] {
	return TypeFieldFunctions[int8]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int8) error {
			encoder.AddInt8(name, value)
			return nil
		},
		IsNonZero: func(i int8) bool {
			return i != 0
		},
	}
}

func Int8(name string, value int8) TypedField[int8] {
	return NewTypedField(
		int8TypeFns(),
		name,
		value,
	)
}
