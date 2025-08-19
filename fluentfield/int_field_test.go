package fluentfield_test

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/fluentfield"
	"go.robertomontagna.dev/zapfluent/util/testing_util"
)

type intTestStruct struct {
	Field1 int
}

func (s intTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(fluentfield.Int("field1", s.Field1).NonZero()).
		Done()
}

func ExampleInt_notEmpty() {
	testing_util.StdOutLogger().Infow(
		"test",
		zap.Object("test_struct", intTestStruct{42}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":42}}
}

func ExampleInt_empty() {
	testing_util.StdOutLogger().Infow(
		"test",
		zap.Object("test_struct", intTestStruct{}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{}}
}

func fpCurrying2to1[P1, P2, R1 any](f func(P1, P2) R1) func(P1) func(P2) R1 {
	return func(p1 P1) func(P2) R1 {
		return func(p2 P2) R1 {
			return f(p1, p2)
		}
	}
}

func ExampleInt_alternative() {
	field := fluentfield.
		Int("field1", 5).
		NonZero().
		Format(fpCurrying2to1(strings.Repeat)("."))

	testing_util.StdOutLogger().Infow(
		"test",
		zap.Object("test_struct", zapcore.ObjectMarshalerFunc(field.Encode)),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":"....."}}
}
