package fluentfield_test

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/util/testing_util"
)

type comparableObjectTestStruct struct {
	Field1 intTestStruct
	Field2 string
}

func (s comparableObjectTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(fluentfield.ComparableObject("field1", s.Field1).NonZero()).
		Add(fluentfield.String("field2", s.Field2).NonZero()).
		Done()
}

func ExampleComparableObject_notEmpty() {
	testing_util.StdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableObjectTestStruct{Field1: intTestStruct{42}}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":{"field1":42}}}
}

func ExampleComparableObject_empty() {
	testing_util.StdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableObjectTestStruct{Field1: intTestStruct{}}),
	)
	// Output:
	//{"level":"info","msg":"test","test_struct":{}}
}
