package lazyoptional_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

func TestMap(t *testing.T) {
	t.Run("on Some", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)
		val, ok := mappedOpt.Get()

		assert.True(t, ok)
		assert.Equal(t, "42", val)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)
		_, ok := mappedOpt.Get()

		assert.False(t, ok)
	})
}
