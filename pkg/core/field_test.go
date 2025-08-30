package core_test

import (
	"strconv"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/pkg/core"

	. "github.com/onsi/gomega"
)

const (
	fieldTestFieldName = "test-field"
	testValueString    = "test-value"
)

func TestTypedField_Name(t *testing.T) {
	g := NewWithT(t)

	field := core.String(fieldTestFieldName, testValueString)

	name := field.Name()

	g.Expect(name).To(Equal(fieldTestFieldName))
}

func TestTypedField_Filtering(t *testing.T) {
	testCases := []struct {
		name       string
		inputField core.TypedField[string]
		assertion  func(g *GomegaWithT, fields map[string]any)
	}{
		{
			name:       "Encode: when value is present, it encodes the value",
			inputField: core.String(fieldTestFieldName, testValueString),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(fieldTestFieldName, testValueString))
			},
		},
		{
			name: "Encode: when value is not present, it does not encode anything",
			inputField: core.String(fieldTestFieldName, testValueString).
				Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name: "Filter: when condition is met, it keeps the value",
			inputField: core.String(fieldTestFieldName, testValueString).
				Filter(func(s string) bool { return true }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name: "Filter: when condition is not met, it removes the value",
			inputField: core.String(fieldTestFieldName, testValueString).
				Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is not zero, it keeps the value",
			inputField: core.String(fieldTestFieldName, testValueString).NonZero(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).ToNot(BeEmpty())
			},
		},
		{
			name:       "NonZero: when value is zero, it removes the value",
			inputField: core.String(fieldTestFieldName, "").NonZero(),
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

	field := core.Int(fieldTestFieldName, 5)

	formattedField := field.Format(func(i int) string { return "formatted-" + strconv.Itoa(i) })
	enc := zapcore.NewMapObjectEncoder()

	err := formattedField.Encode(enc)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(enc.Fields).To(HaveKeyWithValue(fieldTestFieldName, "formatted-5"))
}

func TestTypedPointerField_Name(t *testing.T) {
	g := NewWithT(t)
	nonNilValue := "value"

	field := core.StringPtr(fieldTestFieldName, &nonNilValue)

	name := field.Name()

	g.Expect(name).To(Equal(fieldTestFieldName))
}

func TestTypedPointerField_Encode(t *testing.T) {
	nonNilValue := testValueString
	var nonNilPtr = &nonNilValue

	testCases := []struct {
		name      string
		field     core.Field
		assertion func(g *GomegaWithT, fields map[string]any)
	}{
		{
			name:  "when pointer is not nil, it encodes the value",
			field: core.StringPtr(fieldTestFieldName, nonNilPtr),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(fieldTestFieldName, testValueString))
			},
		},
		{
			name:  "when pointer is nil, it encodes '<nil>'",
			field: core.StringPtr(fieldTestFieldName, nil),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(fieldTestFieldName, "<nil>"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			tc.assertion(g, enc.Fields)
		})
	}
}

func TestTypedPointerField_NonNil(t *testing.T) {
	nonNilValue := testValueString
	var nonNilPtr = &nonNilValue

	testCases := []struct {
		name      string
		field     core.Field
		assertion func(g *GomegaWithT, fields map[string]any)
	}{
		{
			name:  "when pointer is not nil, it returns a valid field",
			field: core.StringPtr(fieldTestFieldName, nonNilPtr).NonNil(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(HaveKeyWithValue(fieldTestFieldName, testValueString))
			},
		},
		{
			name:  "when pointer is nil, it returns an empty field",
			field: core.StringPtr(fieldTestFieldName, nil).NonNil(),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name: "allows chaining when pointer is not nil",
			field: core.StringPtr(fieldTestFieldName, nonNilPtr).NonNil().
				Filter(func(s string) bool { return false }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
		{
			name: "chaining is a no-op when pointer is nil",
			field: core.StringPtr(fieldTestFieldName, nil).NonNil().
				Filter(func(s string) bool { return true }),
			assertion: func(g *GomegaWithT, fields map[string]any) {
				g.Expect(fields).To(BeEmpty())
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			enc := zapcore.NewMapObjectEncoder()

			err := tc.field.Encode(enc)

			g.Expect(err).ToNot(HaveOccurred())
			tc.assertion(g, enc.Fields)
		})
	}
}
