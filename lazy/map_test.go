package lazy_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/lazy"
)

func TestMap(t *testing.T) {
	t.Run("on Some", func(t *testing.T) {
		opt := lazy.Some(42)
		mappedOpt := lazy.Map(opt, func(i int) string {
			return strconv.Itoa(i)
		})
		val, ok := mappedOpt.Get()
		assert.True(t, ok)
		assert.Equal(t, "42", val)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazy.Empty[int]()
		mappedOpt := lazy.Map(opt, func(i int) string {
			return strconv.Itoa(i)
		})
		_, ok := mappedOpt.Get()
		assert.False(t, ok)
	})
}
