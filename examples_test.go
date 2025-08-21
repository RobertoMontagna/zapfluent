package zapfluent_test

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// stdOutLogger creates a new zap.SugaredLogger that is configured to write
// JSON-formatted logs to standard output.
func stdOutLogger() *zap.SugaredLogger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	coreEncoder := zapcore.NewCore(
		core.NewFluentEncoder(
			zapcore.NewJSONEncoder(encoderCfg),
			core.NewConfiguration(),
		),
		os.Stdout,
		zap.DebugLevel,
	)
	logger := zap.New(coreEncoder)
	zap.ReplaceGlobals(logger)
	return zap.S()
}

// test structs
type comparableObjectTestStruct struct {
	Field1 intTestStruct
	Field2 string
}

func (s comparableObjectTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.ComparableObject("field1", s.Field1).NonZero()).
		Add(zapfluent.String("field2", s.Field2).NonZero()).
		Done()
}

type int8TestStruct struct {
	Field1 int8
}

func (s int8TestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.Int8("field1", s.Field1).NonZero()).
		Done()
}

type intTestStruct struct {
	Field1 int
}

func (s intTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.Int("field1", s.Field1).NonZero()).
		Done()
}

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
		Add(zapfluent.Object("field1", s.Field1, core.ReflectiveIsNotNil).NonZero()).
		Done()
}

type stringTestStruct struct {
	Field1 string
}

func (s stringTestStruct) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.String("field1", s.Field1).NonZero()).
		Done()
}

// example tests

func ExampleComparableObject_notEmpty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableObjectTestStruct{Field1: intTestStruct{42}}),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":{"field1":42}}}
}

func ExampleComparableObject_empty() {
	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", comparableObjectTestStruct{Field1: intTestStruct{}}),
	)
	// Output:
	//{"level":"info","msg":"test","test_struct":{}}
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
	field := zapfluent.Int("field1", 5).
		NonZero().
		Format(fpCurrying2to1(strings.Repeat)("."))

	stdOutLogger().Infow(
		"test",
		zap.Object("test_struct", zapcore.ObjectMarshalerFunc(field.Encode)),
	)
	// Output: {"level":"info","msg":"test","test_struct":{"field1":"....."}}
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
