package zapfluent_test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

type stringTestStruct struct {
	Field1 string
}

func (s stringTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.String("field1", s.Field1).NonZero()).
		Done()
}

func ExampleString_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", stringTestStruct{"test"}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":"test"}}
}

func ExampleString_empty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", stringTestStruct{}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{}}
}
