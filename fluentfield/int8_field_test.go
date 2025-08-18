package fluentfield_test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

type int8TestStruct struct {
	Field1 int8
}

func (s int8TestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(fluentfield.Int8("field1", s.Field1).NonZero()).
		Done()
}

func ExampleInt8_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", int8TestStruct{42}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":42}}
}

func ExampleInt8_empty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", int8TestStruct{}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{}}
}
