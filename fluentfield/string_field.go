package fluentfield

import "go.uber.org/zap/zapcore"

var (
	stringTypeFns = TypeFieldFunctions[string]{
		EncodeFunc: func(encoder zapcore.ObjectEncoder, name string, value string) error {
			encoder.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool {
			return s != ""
		},
	}
)

func String(name string, value string) TypedField[string] {
	return NewTypedField(
		stringTypeFns,
		name,
		value,
	)
}
