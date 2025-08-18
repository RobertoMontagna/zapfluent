package fluentfield

import "go.uber.org/zap/zapcore"

func intTypeFns() TypeFieldFunctions[int] {
	return TypeFieldFunctions[int]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value int) error {
			encoder.AddInt(name, value)
			return nil
		},
		IsNonZero: func(i int) bool {
			return i != 0
		},
	}
}

func Int(name string, value int) TypedField[int] {
	return NewTypedField(
		intTypeFns(),
		name,
		value,
	)
}
