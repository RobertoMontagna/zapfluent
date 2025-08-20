package enum_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"go.robertomontagna.dev/zapfluent/enum"
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
	g := NewWithT(t)

	t.Run("with known value", func(t *testing.T) {
		val := testEnumValue1
		s := testEnumSpec.String(val)
		g.Expect(s).To(Equal("Value1"))
	})

	t.Run("with unknown value", func(t *testing.T) {
		val := testEnum(99)
		s := testEnumSpec.String(val)
		g.Expect(s).To(Equal("Unknown(99)"))
	})
}

func TestEnum_FromInt(t *testing.T) {
	g := NewWithT(t)

	t.Run("with valid int", func(t *testing.T) {
		i := 1
		val := testEnumSpec.FromInt(i)
		g.Expect(val).To(Equal(testEnumValue1))
	})

	t.Run("with invalid int", func(t *testing.T) {
		i := 99
		val := testEnumSpec.FromInt(i)
		g.Expect(val).To(Equal(testEnumUnknown))
	})
}
