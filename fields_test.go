package zapfluent_test

import (
	"fmt"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/testutil"
	"go.uber.org/zap/zapcore"

	. "github.com/onsi/gomega"
	"go.robertomontagna.dev/zapfluent"
	"go.robertomontagna.dev/zapfluent/pkg/core"
)

// testObject is a simple struct that implements zapcore.ObjectMarshaler.
type testObject struct {
	Value string
}

// MarshalLogObject implements the zapcore.ObjectMarshaler interface.
func (t testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return zapfluent.AsFluent(enc).
		Add(zapfluent.String("value", t.Value)).
		Done()
}

func TestTypedPointerField_WithAddress_ShouldEncodeValueAndAddress(t *testing.T) {
	strVal := "value"
	intVal := 42

	testCases := []struct {
		name          string
		field         zapfluent.Field
		expectedKey   string
		expectedValue map[string]interface{}
	}{
		{
			name:        "with string pointer (non-nil)",
			field:       zapfluent.StringPtr("my_string_ptr", &strVal).WithAddress(),
			expectedKey: "my_string_ptr",
			expectedValue: map[string]interface{}{
				"value":   "value",
				"address": fmt.Sprintf("%p", &strVal),
			},
		},
		{
			name:        "with string pointer (nil)",
			field:       zapfluent.StringPtr("my_string_ptr", nil).WithAddress(),
			expectedKey: "my_string_ptr",
			expectedValue: map[string]interface{}{
				"value":   "<nil>",
				"address": "0x0",
			},
		},
		{
			name:        "with int pointer (non-nil)",
			field:       zapfluent.IntPtr("my_int_ptr", &intVal).WithAddress(),
			expectedKey: "my_int_ptr",
			expectedValue: map[string]interface{}{
				"value":   intVal,
				"address": fmt.Sprintf("%p", &intVal),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			fluent := core.AsFluent(
				core.NewFluentEncoder(
					testutil.NewDoNotEncodeEncoderForTest(enc),
					core.NewConfiguration(),
				),
			)

			err := fluent.Add(tc.field).Done()

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
		})
	}
}

func TestFluent_Add_ForDifferentFieldTypes_ShouldEncodeCorrectly(t *testing.T) {
	strVal := "value"
	intVal := 42
	int8Val := int8(8)
	boolTrue := true

	testCases := []struct {
		name          string
		field         zapfluent.Field
		expectedKey   string
		expectedValue any
	}{
		{
			name:          "with string",
			field:         zapfluent.String("my_string", "value"),
			expectedKey:   "my_string",
			expectedValue: "value",
		},
		{
			name:          "with string pointer (non-nil)",
			field:         zapfluent.StringPtr("my_string_ptr", &strVal),
			expectedKey:   "my_string_ptr",
			expectedValue: "value",
		},
		{
			name:          "with string pointer (nil)",
			field:         zapfluent.StringPtr("my_string_ptr", nil),
			expectedKey:   "my_string_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with int",
			field:         zapfluent.Int("my_int", 123),
			expectedKey:   "my_int",
			expectedValue: 123,
		},
		{
			name:          "with int pointer (non-nil)",
			field:         zapfluent.IntPtr("my_int_ptr", &intVal),
			expectedKey:   "my_int_ptr",
			expectedValue: 42,
		},
		{
			name:          "with int pointer (nil)",
			field:         zapfluent.IntPtr("my_int_ptr", nil),
			expectedKey:   "my_int_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with int8",
			field:         zapfluent.Int8("my_int8", 12),
			expectedKey:   "my_int8",
			expectedValue: int8(12),
		},
		{
			name:          "with int8 pointer (non-nil)",
			field:         zapfluent.Int8Ptr("my_int8_ptr", &int8Val),
			expectedKey:   "my_int8_ptr",
			expectedValue: int8(8),
		},
		{
			name:          "with int8 pointer (nil)",
			field:         zapfluent.Int8Ptr("my_int8_ptr", nil),
			expectedKey:   "my_int8_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with bool (true)",
			field:         zapfluent.Bool("my_bool", true),
			expectedKey:   "my_bool",
			expectedValue: true,
		},
		{
			name:          "with bool pointer (non-nil)",
			field:         zapfluent.BoolPtr("my_bool_ptr", &boolTrue),
			expectedKey:   "my_bool_ptr",
			expectedValue: true,
		},
		{
			name:          "with bool pointer (nil)",
			field:         zapfluent.BoolPtr("my_bool_ptr", nil),
			expectedKey:   "my_bool_ptr",
			expectedValue: "<nil>",
		},
		{
			name:          "with bool (false)",
			field:         zapfluent.Bool("my_bool", false),
			expectedKey:   "my_bool",
			expectedValue: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			fluent := core.AsFluent(
				core.NewFluentEncoder(
					testutil.NewDoNotEncodeEncoderForTest(enc),
					core.NewConfiguration(),
				),
			)

			err := fluent.Add(tc.field).Done()

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
		})
	}
}

func TestFluent_Add_ForComparableObjectFields_ShouldEncodeCorrectly(t *testing.T) {
	nonZeroObj := &comparableTestObject{Value: "non-zero"}
	zeroObj := &comparableTestObject{Value: ""}

	testCases := []struct {
		name           string
		field          zapfluent.Field
		expectedKey    string
		expectedValue  any
		expectOmission bool
	}{
		{
			name: "with object",
			field: zapfluent.ComparableObject(
				"my_object",
				*nonZeroObj,
			),
			expectedKey: "my_object",
			expectedValue: map[string]interface{}{
				"value": "non-zero",
			},
		},
		{
			name: "with object pointer (non-nil)",
			field: zapfluent.ComparableObjectPtr(
				"my_object_ptr",
				nonZeroObj,
			),
			expectedKey: "my_object_ptr",
			expectedValue: map[string]interface{}{
				"value": "non-zero",
			},
		},
		{
			name: "with object pointer (nil)",
			field: zapfluent.ComparableObjectPtr(
				"my_object_ptr",
				(*comparableTestObject)(nil),
			),
			expectedKey:   "my_object_ptr",
			expectedValue: "<nil>",
		},
		{
			name: "with object (zero value)",
			field: zapfluent.ComparableObject(
				"my_object",
				*zeroObj,
			).NonZero(),
			expectOmission: true,
		},
		{
			name: "with object pointer (zero value)",
			field: zapfluent.ComparableObjectPtr(
				"my_object_ptr",
				zeroObj,
			).NonNil().NonZero(),
			expectOmission: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			fluent := core.AsFluent(
				core.NewFluentEncoder(
					testutil.NewDoNotEncodeEncoderForTest(enc),
					core.NewConfiguration(),
				),
			)

			err := fluent.Add(tc.field).Done()

			g.Expect(err).ToNot(HaveOccurred())

			if tc.expectOmission {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
			}
		})
	}
}

func TestFluent_Add_ForObjectFields_ShouldEncodeCorrectly(t *testing.T) {
	nonZeroObj := &testObject{Value: "non-zero"}
	zeroObj := &testObject{Value: ""}

	isNonZero := func(o testObject) bool {
		return o.Value != ""
	}

	testCases := []struct {
		name           string
		field          zapfluent.Field
		expectedKey    string
		expectedValue  any
		expectOmission bool
	}{
		{
			name: "with object",
			field: zapfluent.Object(
				"my_object",
				*nonZeroObj,
				isNonZero,
			),
			expectedKey: "my_object",
			expectedValue: map[string]interface{}{
				"value": "non-zero",
			},
		},
		{
			name: "with object pointer (non-nil)",
			field: zapfluent.ObjectPtr(
				"my_object_ptr",
				nonZeroObj,
				isNonZero,
			),
			expectedKey: "my_object_ptr",
			expectedValue: map[string]interface{}{
				"value": "non-zero",
			},
		},
		{
			name: "with object pointer (nil)",
			field: zapfluent.ObjectPtr(
				"my_object_ptr",
				(*testObject)(nil),
				isNonZero,
			),
			expectedKey:   "my_object_ptr",
			expectedValue: "<nil>",
		},
		{
			name: "with object (zero value)",
			field: zapfluent.Object(
				"my_object",
				*zeroObj,
				isNonZero,
			).NonZero(),
			expectOmission: true,
		},
		{
			name: "with object pointer (zero value)",
			field: zapfluent.ObjectPtr(
				"my_object_ptr",
				zeroObj,
				isNonZero,
			).NonNil().NonZero(),
			expectOmission: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			fluent := core.AsFluent(
				core.NewFluentEncoder(
					testutil.NewDoNotEncodeEncoderForTest(enc),
					core.NewConfiguration(),
				),
			)

			err := fluent.Add(tc.field).Done()

			g.Expect(err).ToNot(HaveOccurred())

			if tc.expectOmission {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
			}
		})
	}
}

// comparableTestObject is a simple struct that implements zapcore.ObjectMarshaler
// and is comparable.
type comparableTestObject struct {
	Value string
}

// MarshalLogObject implements the zapcore.ObjectMarshaler interface.
func (t comparableTestObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.Value)
	return nil
}
