package zapfluent_test

import (
	"go.robertomontagna.dev/zapfluent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type comparableStructTestStruct struct {
	Field1 intTestStruct
	Field2 string
}

func (s comparableStructTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.NewFluent(enc).
		Add(zapfluent.ComparableStruct("field1", s.Field1).NonZero()).
		Add(zapfluent.String("field2", s.Field2).NonZero()).
		Done()
}

func ExampleComparableStruct_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableStructTestStruct{Field1: intTestStruct{42}}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":{"field1":42}}}
}

func ExampleComparableStruct_empty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableStructTestStruct{Field1: intTestStruct{}}),
	)
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableStructTestStruct{Field1: intTestStruct{}}),
	)
	// Output:
	//{"level":"info","msg":"test","test_struct":{}}
	//{"level":"info","msg":"test","test_struct":{}}
}
