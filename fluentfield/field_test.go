package fluentfield

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
)

const (
	testFieldName = "test-field"
)

func TestLazyTypedField_Encode(t *testing.T) {
	g := NewWithT(t)

	t.Run("when value is present, it encodes the value", func(t *testing.T) {
		functions := typeFieldFunctions[string]{
			EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
				enc.AddString(name, value)
				return nil
			},
			IsNonZero: func(s string) bool { return s != "" },
		}
		field := newTypedField(functions, testFieldName, "test-value")
		enc := zapcore.NewMapObjectEncoder()

		err := field.Encode(enc)

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(HaveKeyWithValue(testFieldName, "test-value"))
	})

	t.Run("when value is not present, it does not encode anything", func(t *testing.T) {
		functions := typeFieldFunctions[string]{
			EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
				enc.AddString(name, value)
				return nil
			},
			IsNonZero: func(s string) bool { return s != "" },
		}
		field := newTypedField(functions, testFieldName, "test-value").
			Filter(func(s string) bool { return false }) // This will make the value not present
		enc := zapcore.NewMapObjectEncoder()

		err := field.Encode(enc)

		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(enc.Fields).To(BeEmpty())
	})
}

func TestLazyTypedField_Name(t *testing.T) {
	g := NewWithT(t)
	functions := typeFieldFunctions[string]{}
	field := newTypedField(functions, testFieldName, "test-value")

	g.Expect(field.Name()).To(Equal(testFieldName))
}

func TestLazyTypedField_Filter(t *testing.T) {
	g := NewWithT(t)
	functions := typeFieldFunctions[string]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
			enc.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool { return s != "" },
	}
	field := newTypedField(functions, testFieldName, "test-value")

	t.Run("when condition is met, it keeps the value", func(t *testing.T) {
		filteredField := field.Filter(func(s string) bool { return true })
		enc := zapcore.NewMapObjectEncoder()
		_ = filteredField.Encode(enc)
		g.Expect(enc.Fields).ToNot(BeEmpty())
	})

	t.Run("when condition is not met, it removes the value", func(t *testing.T) {
		filteredField := field.Filter(func(s string) bool { return false })
		enc := zapcore.NewMapObjectEncoder()
		_ = filteredField.Encode(enc)
		g.Expect(enc.Fields).To(BeEmpty())
	})
}

func TestLazyTypedField_NonZero(t *testing.T) {
	g := NewWithT(t)
	functions := typeFieldFunctions[string]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value string) error {
			enc.AddString(name, value)
			return nil
		},
		IsNonZero: func(s string) bool { return s != "" },
	}

	t.Run("when value is not zero, it keeps the value", func(t *testing.T) {
		field := newTypedField(functions, testFieldName, "test-value").NonZero()
		enc := zapcore.NewMapObjectEncoder()
		_ = field.Encode(enc)
		g.Expect(enc.Fields).ToNot(BeEmpty())
	})

	t.Run("when value is zero, it removes the value", func(t *testing.T) {
		field := newTypedField(functions, testFieldName, "").NonZero()
		enc := zapcore.NewMapObjectEncoder()
		_ = field.Encode(enc)
		g.Expect(enc.Fields).To(BeEmpty())
	})
}

func TestLazyTypedField_Format(t *testing.T) {
	g := NewWithT(t)
	functions := typeFieldFunctions[int]{
		EncodeFunc: func(enc zapcore.ObjectEncoder, name string, value int) error {
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
		{"nil interface", (interface{})(nil), false},
		{"nil slice", ([]int)(nil), false},
		{"nil map", (map[int]int)(nil), false},
		{"nil channel", (chan int)(nil), false},
		{"nil func", (func())(nil), false},
		{"non-nil int", 1, true},
		{"non-nil string", "hello", true},
		{"non-nil struct", struct{}{}, true},
		{"non-nil pointer", new(int), true},
		{"non-nil interface", interface{}(1), true},
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
