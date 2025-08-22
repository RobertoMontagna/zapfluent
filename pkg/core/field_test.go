package core

import (
	"testing"

	"go.uber.org/zap/zapcore"

	. "github.com/onsi/gomega"
)

const (
	testValueString = "test-value"
)

func TestLazyTypedField_Name(t *testing.T) {
	g := NewWithT(t)

	functions := typeFieldFunctions[string]{}
	field := newTypedField(functions, testFieldName, testValueString)

	g.Expect(field.Name()).To(Equal(testFieldName))
}

func TestLazyTypedField(t *testing.T) {
	functions := typeFieldFunctions[string]{
		encodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
			enc.AddString(name, value)
			return nil
		},
		isNonZero: func(s string) bool { return s != "" },
	}
	baseField := newTypedField(functions, testFieldName, testValueString)
	zeroField := newTypedField(functions, testFieldName, "")

	testCases := []struct {
		name       string
		inputField Field
		assertion  func(g *GomegaWithT, fields map[string]any)
	}{
		{
			name:       "Encode: when value is present, it encodes the value",
			inputField: baseField,
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(testFieldName, testValueString))
			},
		},
		{
			name:       "Encode: when value is not present, it does not encode anything",
			inputField: baseField.Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name:       "Filter: when condition is met, it keeps the value",
			inputField: baseField.Filter(func(s string) bool { return true }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name:       "Filter: when condition is not met, it removes the value",
			inputField: baseField.Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is not zero, it keeps the value",
			inputField: baseField.NonZero(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is zero, it removes the value",
			inputField: zeroField.NonZero(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			enc := zapcore.NewMapObjectEncoder()

			err := tc.inputField.Encode(enc)
			g.Expect(err).ToNot(HaveOccurred())

			tc.assertion(g, enc.Fields)
		})
	}
}

func TestLazyTypedField_Format(t *testing.T) {
	g := NewWithT(t)

	functions := typeFieldFunctions[int]{
		encodeFunc: func(enc zapcore.ObjectEncoder, name string, value int) error {
			enc.AddInt(name, value)
			return nil
		},
	}
	field := newTypedField(functions, testFieldName, 5)
	formattedField := field.Format(func(i int) string { return "formatted" })
	enc := zapcore.NewMapObjectEncoder()

	err := formattedField.Encode(enc)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, "formatted"))
}

func TestIsNotNil_WithUntypedNil(t *testing.T) {
	g := NewWithT(t)

	var input any

	actual := ReflectiveIsNotNil(input)

	g.Expect(actual).To(BeFalse())
}

func TestIsNotNil_WithTypedValues(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		name     string
		input    any
		expected bool
	}{
		{"nil pointer", (*int)(nil), false},
		{"nil interface", (any)(nil), false},
		{"nil slice", ([]int)(nil), false},
		{"nil map", (map[int]int)(nil), false},
		{"nil channel", (chan int)(nil), false},
		{"nil func", (func())(nil), false},
		{"non-nil int", 1, true},
		{"non-nil string", "hello", true},
		{"non-nil struct", struct{}{}, true},
		{"non-nil pointer", new(int), true},
		{"non-nil interface", any(1), true},
		{"non-nil slice", make([]int, 1), true},
		{"non-nil map", make(map[int]int), true},
		{"non-nil channel", make(chan int), true},
		{"non-nil func", func() {}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ReflectiveIsNotNil(tc.input)

			g.Expect(actual).To(Equal(tc.expected))
		})
	}
}
