package zapfluent_test

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
)

type intTestStruct struct {
	Field1 int
}

func (s intTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.Int(s.Field1).NonZero().Name("field1")).
		Done()
}

func ExampleInt_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", intTestStruct{42}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":42}}
}

func ExampleInt_empty() {
	stdOutLogger().Infow(
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
	alternative := func(s intTestStruct) zapcore.ObjectMarshalerFunc {
		return zapfluent.
			Int(s.Field1).
			NonZero().
			Format(fpCurrying2to1(strings.Repeat)(".")).
			Name("field1")
	}

	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", alternative(intTestStruct{5})),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":"....."}}
}
