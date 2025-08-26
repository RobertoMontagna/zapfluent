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

func (t *testObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
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
	isNonZero := func(o *testObject) bool { return o != nil && o.value != "" }

	testCases := []struct {
		name          string
		field         core.TypedField[*testObject]
		expectedKey   string
		expectedValue any
		shouldBeEmpty bool
	}{
		{
			name:        "it creates an object field correctly",
			field:       core.Object("my-key", &testObject{value: "test"}, isNonZero),
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
				&testObject{value: "value"},
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
			field:         core.Object("zero-key", &testObject{value: ""}, isNonZero).NonZero(),
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
