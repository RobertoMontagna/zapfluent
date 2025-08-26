package enum_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/enum"

	. "github.com/onsi/gomega"
)

type testEnum int

const (
	testEnumUnknown testEnum = iota
	testEnumValue1
	testEnumValue2
)

var testEnumSpec = enum.New(
	map[testEnum]string{
		testEnumUnknown: "Unknown",
		testEnumValue1:  "Value1",
		testEnumValue2:  "Value2",
	},
	testEnumUnknown,
)

func TestEnum_String(t *testing.T) {
	testCases := []struct {
		name     string
		value    testEnum
		expected string
	}{
		{
			name:     "with known value",
			value:    testEnumValue1,
			expected: "Value1",
		},
		{
			name:     "with unknown value",
			value:    testEnum(99),
			expected: "Unknown(99)",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			s := testEnumSpec.String(tc.value)

			g.Expect(s).To(Equal(tc.expected))
		})
	}
}

func TestEnum_FromInt(t *testing.T) {
	testCases := []struct {
		name     string
		value    int
		expected testEnum
	}{
		{
			name:     "with valid int",
			value:    1,
			expected: testEnumValue1,
		},
		{
			name:     "with invalid int",
			value:    99,
			expected: testEnumUnknown,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			val := testEnumSpec.FromInt(tc.value)

			g.Expect(val).To(Equal(tc.expected))
		})
	}
}
