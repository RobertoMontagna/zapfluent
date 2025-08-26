package core_test

import (
	"strconv"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

const (
	testValueString = "test-value"
)

func TestTypedField_Name(t *testing.T) {
	g := NewWithT(t)

	field := core.String(testFieldName, testValueString)

	g.Expect(field.Name()).To(Equal(testFieldName))
}

func TestTypedField_Filtering(t *testing.T) {
	testCases := []struct {
		name       string
		inputField core.TypedField[string]
		assertion  func(g *GomegaWithT, fields map[string]any)
	}{
		{
			name:       "Encode: when value is present, it encodes the value",
			inputField: core.String(testFieldName, testValueString),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(testFieldName, testValueString))
			},
		},
		{
			name: "Encode: when value is not present, it does not encode anything",
			inputField: core.String(testFieldName, testValueString).
				Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name: "Filter: when condition is met, it keeps the value",
			inputField: core.String(testFieldName, testValueString).
				Filter(func(s string) bool { return true }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name: "Filter: when condition is not met, it removes the value",
			inputField: core.String(testFieldName, testValueString).
				Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is not zero, it keeps the value",
			inputField: core.String(testFieldName, testValueString).NonZero(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is zero, it removes the value",
			inputField: core.String(testFieldName, "").NonZero(),
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

func TestTypedField_Format(t *testing.T) {
	g := NewWithT(t)

	field := core.Int(testFieldName, 5)

	formattedField := field.Format(func(i int) string { return "formatted-" + strconv.Itoa(i) })
	enc := zapcore.NewMapObjectEncoder()

	err := formattedField.Encode(enc)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, "formatted-5"))
}

func TestIsNotNil_WithUntypedNil(t *testing.T) {
	g := NewWithT(t)

	var input any

	actual := core.ReflectiveIsNotNil(input)

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
			actual := core.ReflectiveIsNotNil(tc.input)

			g.Expect(actual).To(Equal(tc.expected))
		})
	}
}
