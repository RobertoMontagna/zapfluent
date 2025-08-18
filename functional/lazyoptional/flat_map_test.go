package lazyoptional_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

func TestFlatMap(t *testing.T) {
	t.Run("on Some that returns Some", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Some(strconv.Itoa(i)) }

		fmOpt := lazyoptional.FlatMap(opt, f)
		val, ok := fmOpt.Get()

		assert.True(t, ok)
		assert.Equal(t, "42", val)
	})

	t.Run("on Some that returns Empty", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Empty[string]() }

		fmOpt := lazyoptional.FlatMap(opt, f)
		_, ok := fmOpt.Get()

		assert.False(t, ok)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Some(strconv.Itoa(i)) }

		fmOpt := lazyoptional.FlatMap(opt, f)
		_, ok := fmOpt.Get()

		assert.False(t, ok)
	})
}
