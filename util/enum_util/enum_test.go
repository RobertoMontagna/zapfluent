package enum_util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/util/enum_util"
)

type testEnum int

const (
	testEnumUnknown testEnum = iota
	testEnumValue1
	testEnumValue2
)

var testEnumHelper = enum_util.NewUtilEnum(
	map[testEnum]string{
		testEnumUnknown: "Unknown",
		testEnumValue1:  "Value1",
		testEnumValue2:  "Value2",
	},
	testEnumUnknown,
)

func TestUtilEnum_String(t *testing.T) {
	t.Run("with known value", func(t *testing.T) {
		val := testEnumValue1
		s := testEnumHelper.String(val)
		assert.Equal(t, "Value1", s)
	})

	t.Run("with unknown value", func(t *testing.T) {
		val := testEnum(99)
		s := testEnumHelper.String(val)
		assert.Equal(t, "Unknown(99)", s)
	})
}

func TestUtilEnum_FromInt(t *testing.T) {
	t.Run("with valid int", func(t *testing.T) {
		i := 1
		val := testEnumHelper.FromInt(i)
		assert.Equal(t, testEnumValue1, val)
	})

	t.Run("with invalid int", func(t *testing.T) {
		i := 99
		val := testEnumHelper.FromInt(i)
		assert.Equal(t, testEnumUnknown, val)
	})
}
