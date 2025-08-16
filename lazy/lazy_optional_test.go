package lazy_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/lazy"
)

func TestLazyOptional_Some(t *testing.T) {
	opt := lazy.Some(42)
	val, ok := opt.Get()
	assert.True(t, ok)
	assert.Equal(t, 42, val)
}

func TestLazyOptional_Empty(t *testing.T) {
	opt := lazy.Empty[int]()
	_, ok := opt.Get()
	assert.False(t, ok)
}

func TestLazyOptional_Filter(t *testing.T) {
	// Test case 1: Filter passes
	opt1 := lazy.Some(42).Filter(func(i int) bool { return i > 10 })
	val1, ok1 := opt1.Get()
	assert.True(t, ok1)
	assert.Equal(t, 42, val1)

	// Test case 2: Filter fails
	opt2 := lazy.Some(42).Filter(func(i int) bool { return i < 10 })
	_, ok2 := opt2.Get()
	assert.False(t, ok2)

	// Test case 3: Filter on empty
	opt3 := lazy.Empty[int]().Filter(func(i int) bool { return i > 10 })
	_, ok3 := opt3.Get()
	assert.False(t, ok3)
}

func TestLazyOptional_Map(t *testing.T) {
	// Test case 1: Map on Some
	opt1 := lazy.Some(42)
	mappedOpt1 := lazy.Map(opt1, func(i int) string {
		return strconv.Itoa(i)
	})
	val1, ok1 := mappedOpt1.Get()
	assert.True(t, ok1)
	assert.Equal(t, "42", val1)

	// Test case 2: Map on Empty
	opt2 := lazy.Empty[int]()
	mappedOpt2 := lazy.Map(opt2, func(i int) string {
		return strconv.Itoa(i)
	})
	_, ok2 := mappedOpt2.Get()
	assert.False(t, ok2)
}
