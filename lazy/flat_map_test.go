package lazy_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/lazy"
)

func TestFlatMap(t *testing.T) {
	t.Run("on Some that returns Some", func(t *testing.T) {
		opt := lazy.Some(42)
		fmOpt := lazy.FlatMap(opt, func(i int) lazy.LazyOptional[string] {
			return lazy.Some(strconv.Itoa(i))
		})
		val, ok := fmOpt.Get()
		assert.True(t, ok)
		assert.Equal(t, "42", val)
	})

	t.Run("on Some that returns Empty", func(t *testing.T) {
		opt := lazy.Some(42)
		fmOpt := lazy.FlatMap(opt, func(i int) lazy.LazyOptional[string] {
			return lazy.Empty[string]()
		})
		_, ok := fmOpt.Get()
		assert.False(t, ok)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazy.Empty[int]()
		fmOpt := lazy.FlatMap(opt, func(i int) lazy.LazyOptional[string] {
			return lazy.Some(strconv.Itoa(i))
		})
		_, ok := fmOpt.Get()
		assert.False(t, ok)
	})
}
