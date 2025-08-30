package core_test

import (
	"testing"

	"go.uber.org/zap/zapcore"

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

func TestString(t *testing.T) {
	testCases := []struct {
		name          string
		field         core.TypedField[string]
		expectedKey   string
		expectedValue string
		shouldBeEmpty bool
	}{
		{
			name:          "it creates a string field correctly",
			field:         core.String("my-key", "my-value"),
			expectedKey:   "my-key",
			expectedValue: "my-value",
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.String("non-zero-key", "value").NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: "value",
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.String("zero-key", "").NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestInt(t *testing.T) {
	testCases := []struct {
		name          string
		field         core.TypedField[int]
		expectedKey   string
		expectedValue int
		shouldBeEmpty bool
	}{
		{
			name:          "it creates an int field correctly",
			field:         core.Int("my-key", 123),
			expectedKey:   "my-key",
			expectedValue: 123,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int("non-zero-key", 42).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: 42,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int("zero-key", 0).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestInt8(t *testing.T) {
	testCases := []struct {
		name          string
		field         core.TypedField[int8]
		expectedKey   string
		expectedValue int8
		shouldBeEmpty bool
	}{
		{
			name:          "it creates an int8 field correctly",
			field:         core.Int8("my-key", 12),
			expectedKey:   "my-key",
			expectedValue: 12,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with non-zero value",
			field:         core.Int8("non-zero-key", 4).NonZero(),
			expectedKey:   "non-zero-key",
			expectedValue: 4,
			shouldBeEmpty: false,
		},
		{
			name:          "NonZero filter works correctly with zero value",
			field:         core.Int8("zero-key", 0).NonZero(),
			expectedKey:   "zero-key",
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestObject(t *testing.T) {
	isNonZero := func(o testObject) bool { return o.value != "" }

	testCases := []struct {
		name          string
		field         core.TypedField[testObject]
		expectedKey   string
		expectedValue any
		shouldBeEmpty bool
	}{
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestBool(t *testing.T) {
	testCases := []struct {
		name          string
		field         core.TypedField[bool]
		expectedKey   string
		expectedValue bool
		shouldBeEmpty bool
	}{
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestComparableObject(t *testing.T) {
	testCases := []struct {
		name          string
		field         core.TypedField[testComparableObject]
		expectedKey   string
		expectedValue any
		shouldBeEmpty bool
	}{
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).ToNot(HaveKey(tc.expectedKey))
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(tc.expectedKey, tc.expectedValue))
				g.Expect(tc.field.Name()).To(Equal(tc.expectedKey))
			}
		})
	}
}

func TestStringPtr_Encode(t *testing.T) {
	nonNilValue := "my-value"
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.StringPtr("my-key", &nonNilValue),
			expectedValue: "my-value",
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.StringPtr("my-key", nil),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestStringPtr_NonNil(t *testing.T) {
	nonNilValue := "my-value"
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.StringPtr("my-key", &nonNilValue).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.StringPtr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", nonNilValue))
			}
		})
	}
}

func TestIntPtr_Encode(t *testing.T) {
	nonNilValue := 123
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.IntPtr("my-key", &nonNilValue),
			expectedValue: 123,
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.IntPtr("my-key", nil),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestIntPtr_NonNil(t *testing.T) {
	nonNilValue := 123
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.IntPtr("my-key", &nonNilValue).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.IntPtr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", nonNilValue))
			}
		})
	}
}

func TestInt8Ptr_Encode(t *testing.T) {
	nonNilValue := int8(12)
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.Int8Ptr("my-key", &nonNilValue),
			expectedValue: int8(12),
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.Int8Ptr("my-key", nil),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestInt8Ptr_NonNil(t *testing.T) {
	nonNilValue := int8(12)
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.Int8Ptr("my-key", &nonNilValue).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.Int8Ptr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", nonNilValue))
			}
		})
	}
}

func TestBoolPtr_Encode(t *testing.T) {
	nonNilValue := true
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.BoolPtr("my-key", &nonNilValue),
			expectedValue: true,
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.BoolPtr("my-key", nil),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestBoolPtr_NonNil(t *testing.T) {
	nonNilValue := true
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.BoolPtr("my-key", &nonNilValue).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.BoolPtr("my-key", nil).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", nonNilValue))
			}
		})
	}
}

func TestObjectPtr_Encode(t *testing.T) {
	isNonZero := func(o testObject) bool { return o.value != "" }
	nonNilValue := &testObject{value: "test"}
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.ObjectPtr("my-key", nonNilValue, isNonZero),
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.ObjectPtr("my-key", (*testObject)(nil), isNonZero),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestObjectPtr_NonNil(t *testing.T) {
	isNonZero := func(o testObject) bool { return o.value != "" }
	nonNilValue := &testObject{value: "test"}
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.ObjectPtr("my-key", nonNilValue, isNonZero).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name:          "when pointer is nil, it returns an empty field",
			field:         core.ObjectPtr("my-key", (*testObject)(nil), isNonZero).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(
					"my-key",
					map[string]interface{}{"value": "test"},
				))
			}
		})
	}
}

func TestComparableObjectPtr_Encode(t *testing.T) {
	nonNilValue := &testComparableObject{value: "test"}
	testCases := []struct {
		name          string
		field         core.Field
		expectedValue any
	}{
		{
			name:          "when pointer is not nil, it encodes the value",
			field:         core.ComparableObjectPtr("my-key", nonNilValue),
			expectedValue: map[string]interface{}{"value": "test"},
		},
		{
			name:          "when pointer is nil, it encodes <nil>",
			field:         core.ComparableObjectPtr("my-key", (*testComparableObject)(nil)),
			expectedValue: "<nil>",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(enc.Fields).To(HaveKeyWithValue("my-key", tc.expectedValue))
		})
	}
}

func TestComparableObjectPtr_NonNil(t *testing.T) {
	nonNilValue := &testComparableObject{value: "test"}
	testCases := []struct {
		name          string
		field         core.Field
		shouldBeEmpty bool
	}{
		{
			name:          "when pointer is not nil, it returns a valid field",
			field:         core.ComparableObjectPtr("my-key", nonNilValue).NonNil(),
			shouldBeEmpty: false,
		},
		{
			name: "when pointer is nil, it returns an empty field",
			field: core.ComparableObjectPtr(
				"my-key",
				(*testComparableObject)(nil),
			).NonNil(),
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()
			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			if tc.shouldBeEmpty {
				g.Expect(enc.Fields).To(BeEmpty())
			} else {
				g.Expect(enc.Fields).To(HaveKeyWithValue(
					"my-key",
					map[string]interface{}{"value": "test"},
				))
			}
		})
	}
}
