package lazy_test

import (
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
	t.Run("on Some with passing condition", func(t *testing.T) {
		opt := lazy.Some(42).Filter(func(i int) bool { return i > 10 })
		val, ok := opt.Get()
		assert.True(t, ok)
		assert.Equal(t, 42, val)
	})

	t.Run("on Some with failing condition", func(t *testing.T) {
		opt := lazy.Some(42).Filter(func(i int) bool { return i < 10 })
		_, ok := opt.Get()
		assert.False(t, ok)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazy.Empty[int]().Filter(func(i int) bool { return i > 10 })
		_, ok := opt.Get()
		assert.False(t, ok)
	})
}
