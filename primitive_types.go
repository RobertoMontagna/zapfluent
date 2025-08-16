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

func String(name string, value string) TypedField[string] {
	return stringFunc(name, value)
}

func Int(name string, value int) TypedField[int] {
	return Primitive[int]()(name, value)
}

func Int8(name string, value int8) TypedField[int8] {
	return Primitive[int8]()(name, value)
}
