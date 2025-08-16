package zapfluent

import "go.uber.org/zap/zapcore"

func stringTypeFns() TypeFieldFunctions[string] {
	return TypeFieldFunctions[string]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value string) error {
			encoder.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool {
			return s != ""
		},
		FieldNoop: typedFieldNoop[string]{},
	}
}

func String(name string, value string) TypedField[string] {
	return NewTypedField(
		stringTypeFns(),
		name,
		value,
	)
}
