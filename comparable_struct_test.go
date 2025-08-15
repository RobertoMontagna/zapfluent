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
		Add(zapfluent.ComparableStruct(s.Field1).NonZero().AsPII().Name("field1")).
		Add(zapfluent.String(s.Field2).NonZero().Name("field2")).
		Done()
}

func ExampleComparableStruct_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableStructTestStruct{Field1: intTestStruct{42}}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1 <<pii>>":{"field1":42}}}
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
