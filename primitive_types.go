package zapfluent

import "go.uber.org/zap/zapcore"

var stringFunc = Primitive[string](
	func(t *TypeFieldFunctions[string]) {
		t.EncodeFunc = func(encoder zapcore.ObjectEncoder, name string, value string) error {
			encoder.AddString(name, value)
			return nil
		}
	},
)

func String(value string) TypedField[string] {
	return stringFunc(value)
}

func Int(value int) TypedField[int] {
	return Primitive[int]()(value)
}

func Int8(value int8) TypedField[int8] {
	return Primitive[int8]()(value)
}
