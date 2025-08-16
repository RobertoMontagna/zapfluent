package zapfluent

import "go.uber.org/zap/zapcore"

func String(name string, value string) TypedField[string] {
	return NewTypedField(
		TypeFieldFunctions[string]{
			EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value string) error {
				encoder.AddString(name, value)
				return nil
			},
			IsNonZero: func(s string) bool {
				return s != ""
			},
			FieldNoop: typedFieldNoop[string]{},
		},
		name,
		value,
	)
}
