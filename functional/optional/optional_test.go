package optional_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/functional/optional"
)

func TestOptional_Some(t *testing.T) {
	o := optional.Some("test")
	assert.True(t, o.IsPresent())
	val, ok := o.Get()
	assert.True(t, ok)
	assert.Equal(t, "test", val)
}

func TestOptional_Empty(t *testing.T) {
	o := optional.Empty[string]()
	assert.False(t, o.IsPresent())
	val, ok := o.Get()
	assert.False(t, ok)
	assert.Equal(t, "", val) // Zero value
}

func TestOptional_ForEach(t *testing.T) {
	t.Run("with present value", func(t *testing.T) {
		o := optional.Some("test")
		var result string
		o.ForEach(func(s string) {
			result = s
		})
		assert.Equal(t, "test", result)
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[string]()
		var result string
		o.ForEach(func(s string) {
			result = "should not be called"
		})
		assert.Equal(t, "", result)
	})
}

func TestOptional_Map(t *testing.T) {
	t.Run("with present value", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.Map(o, strconv.Itoa)
		assert.True(t, mapped.IsPresent())
		val, _ := mapped.Get()
		assert.Equal(t, "123", val)
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.Map(o, strconv.Itoa)
		assert.False(t, mapped.IsPresent())
	})
}

func TestOptional_FlatMap(t *testing.T) {
	t.Run("with present value mapping to present", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		assert.True(t, mapped.IsPresent())
		val, _ := mapped.Get()
		assert.Equal(t, "123", val)
	})

	t.Run("with present value mapping to empty", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Empty[string]()
		})
		assert.False(t, mapped.IsPresent())
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		assert.False(t, mapped.IsPresent())
	})
}
