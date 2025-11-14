package core_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/lang"
	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

type testObject struct {
	value string
}

func (t testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.value)
	return nil
}

type testComparableObject struct {
	value string
}

func (t testComparableObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("value", t.value)
	return nil
}

type fieldTestCase struct {
	name          string
	field         core.Field
	shouldBeEmpty bool
	expectedKey   string
	expectedValue any
}

func TestBool(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates a bool field correctly",
			field:         core.Bool("my-key", true),
			expectedKey:   "my-key",
			expectedValue: true,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Bool("non-zero-key", true).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: true,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Bool("zero-key", false).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt32(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates an int32 field correctly",
			field:         core.Int32("my-key", 12),
			expectedKey:   "my-key",
			expectedValue: int32(12),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int32("non-zero-key", 4).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: int32(4),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int32("zero-key", 0).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt32Ptr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.Int32Ptr("my-key", lang.ToPtr(int32(12))),
			expectedKey:   "my-key",
			expectedValue: int32(12),
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.Int32Ptr("my-key", nil),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.Int32Ptr("my-key", lang.ToPtr(int32(12))).NonNil(),
			expectedKey:   "my-key",
			expectedValue: int32(12),
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.Int32Ptr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt16(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates an int16 field correctly",
			field:         core.Int16("my-key", 12),
			expectedKey:   "my-key",
			expectedValue: int16(12),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int16("non-zero-key", 4).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: int16(4),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int16("zero-key", 0).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt16Ptr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.Int16Ptr("my-key", lang.ToPtr(int16(12))),
			expectedKey:   "my-key",
			expectedValue: int16(12),
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.Int16Ptr("my-key", nil),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.Int16Ptr("my-key", lang.ToPtr(int16(12))).NonNil(),
			expectedKey:   "my-key",
			expectedValue: int16(12),
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.Int16Ptr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestBoolPtr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.BoolPtr("my-key", lang.ToPtr(true)),
			expectedKey:   "my-key",
			expectedValue: true,
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.BoolPtr("my-key", nil),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.BoolPtr("my-key", lang.ToPtr(true)).NonNil(),
			expectedKey:   "my-key",
			shouldBeEmpty: false,
			expectedValue: true,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.BoolPtr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestString(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates a string field correctly",
			field:         core.String("my-key", "my-value"),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: "my-value",
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.String("non-zero-key", "value").NonZero(),
			shouldBeEmpty: false,
			expectedKey:   "non-zero-key",
			expectedValue: "value",
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.String("zero-key", "").NonZero(),
			shouldBeEmpty: true,
			expectedKey:   "zero-key",
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestStringPtr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.StringPtr("my-key", lang.ToPtr("my-value")),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: "my-value",
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.StringPtr("my-key", nil),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.StringPtr("my-key", lang.ToPtr("my-value")).NonNil(),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: "my-value",
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.StringPtr("my-key", nil).NonNil(),
			expectedKey:   "my-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates an int field correctly",
			field:         core.Int("my-key", 123),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: 123,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int("non-zero-key", 42).NonZero(),
			shouldBeEmpty: false,
			expectedKey:   "non-zero-key",
			expectedValue: 42,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int("zero-key", 0).NonZero(),
			shouldBeEmpty: true,
			expectedKey:   "zero-key",
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestIntPtr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.IntPtr("my-key", lang.ToPtr(123)),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: 123,
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.IntPtr("my-key", nil),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.IntPtr("my-key", lang.ToPtr(123)),
			shouldBeEmpty: false,
			expectedKey:   "my-key",
			expectedValue: 123,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.IntPtr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
			expectedKey:   "my-key",
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt8(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "it creates an int8 field correctly",
			field:         core.Int8("my-key", 12),
			expectedKey:   "my-key",
			expectedValue: int8(12),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int8("non-zero-key", 4).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: int8(4),
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int8("zero-key", 0).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestInt8Ptr(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.Int8Ptr("my-key", lang.ToPtr(int8(12))),
			expectedKey:   "my-key",
			expectedValue: int8(12),
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.Int8Ptr("my-key", nil),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.Int8Ptr("my-key", lang.ToPtr(int8(12))).NonNil(),
			expectedKey:   "my-key",
			expectedValue: int8(12),
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.Int8Ptr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestObject(t *testing.T) {
	isNonZero := func(o testObject) bool { return o.value != "" }

	testCases := []fieldTestCase{
		{
			name:        "it creates an object field correctly",
			field:       core.Object("my-key", testObject{value: "test"}, isNonZero),
			expectedKey: "my-key",
			expectedValue: map[string]interface{}{
				"value": "test",
			},
			shouldBeEmpty: false,
		},
		{
			name: "NonZero filter works correctly with non-zero value",
			field: core.Object(
				"non-zero-key",
				testObject{value: "value"},
				isNonZero,
			).NonZero(),
			expectedKey: "non-zero-key",
			expectedValue: map[string]interface{}{
				"value": "value",
			},
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Object("zero-key", testObject{value: ""}, isNonZero).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestObjectPtr(t *testing.T) {
	isNonZero := func(o testObject) bool { return o.value != "" }
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.ObjectPtr("my-key", lang.ToPtr(testObject{value: "test"}), isNonZero),
			expectedKey:   "my-key",
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.ObjectPtr("my-key", nil, isNonZero),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.ObjectPtr("my-key", lang.ToPtr(testObject{value: "test"}), isNonZero).NonNil(),
			expectedKey:   "my-key",
			shouldBeEmpty: false,
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.ObjectPtr("my-key", nil, isNonZero).NonNil(),
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestComparableObject(t *testing.T) {
	testCases := []fieldTestCase{
		{
			name:        "it creates a comparable object field correctly",
			field:       core.ComparableObject("my-key", testComparableObject{value: "test"}),
			expectedKey: "my-key",
			expectedValue: map[string]interface{}{
				"value": "test",
			},
			shouldBeEmpty: false,
		},
		{
			name: "NonZero filter works correctly with non-zero value",
			field: core.ComparableObject(
				"non-zero-key",
				testComparableObject{value: "value"},
			).NonZero(),
			expectedKey: "non-zero-key",
			expectedValue: map[string]interface{}{
				"value": "value",
			},
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.ComparableObject("zero-key", testComparableObject{}).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func TestComparableObjectPtr(t *testing.T) {
	nonNilValue := &testComparableObject{value: "test"}
	testCases := []fieldTestCase{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.ComparableObjectPtr("my-key", nonNilValue),
			expectedKey:   "my-key",
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.ComparableObjectPtr("my-key", (*testComparableObject)(nil)),
			expectedKey:   "my-key",
			expectedValue: core.NilSentinel,
		},
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.ComparableObjectPtr("my-key", nonNilValue).NonNil().NonZero(),
			expectedKey:   "my-key",
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it encodes NilSentinel",
			field:         core.ComparableObjectPtr("my-key", (*testComparableObject)(nil)).NonNil(),
			expectedKey:   "my-key",
			shouldBeEmpty: true,
		},
	}

	fieldsTestCaseValidation(t, testCases)
}

func fieldsTestCaseValidation(t *testing.T, testCases []fieldTestCase) {
	t.Helper()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := testCase.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if testCase.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(testCase.expectedKey))
			} else {
				g.Expect(enc.Fields).
					To(HaveKeyWithValue(testCase.expectedKey, testCase.expectedValue))
				g.Expect(testCase.field.Name()).To(Equal(testCase.expectedKey))
			}
		})
	}
}
