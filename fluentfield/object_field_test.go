package fluentfield_test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/fluentfield"
)

// A simple struct that implements zapcore.ObjectMarshaler
type testObject struct {
	value string
}

func (t *testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.value)
	return nil
}

type objectTestStruct struct {
	Field1 *testObject
}

func (s objectTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(fluentfield.Object("field1", s.Field1, fluentfield.IsNotNil).NonZero()).
		Done()
}

func ExampleObject_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", objectTestStruct{Field1: &testObject{"hello"}}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":{"value":"hello"}}}
}

func ExampleObject_empty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", objectTestStruct{Field1: nil}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{}}
}
