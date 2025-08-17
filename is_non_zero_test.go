package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent"
)

func TestIsNotNil(t *testing.T) {
	t.Run("with nil values", func(t *testing.T) {
		var p *int
		var i interface{}
		var s []int
		var m map[int]int
		var c chan int
		var f func()

		assert.False(t, zapfluent.IsNotNil[any](nil))
		assert.False(t, zapfluent.IsNotNil(p))
		assert.False(t, zapfluent.IsNotNil(i))
		assert.False(t, zapfluent.IsNotNil(s))
		assert.False(t, zapfluent.IsNotNil(m))
		assert.False(t, zapfluent.IsNotNil(c))
		assert.False(t, zapfluent.IsNotNil(f))
	})

	t.Run("with non-nil values", func(t *testing.T) {
		p := new(int)
		var i interface{} = 1
		s := make([]int, 1)
		m := make(map[int]int)
		c := make(chan int)
		f := func() {}

		assert.True(t, zapfluent.IsNotNil(1))
		assert.True(t, zapfluent.IsNotNil("hello"))
		assert.True(t, zapfluent.IsNotNil(struct{}{}))
		assert.True(t, zapfluent.IsNotNil(p))
		assert.True(t, zapfluent.IsNotNil(i))
		assert.True(t, zapfluent.IsNotNil(s))
		assert.True(t, zapfluent.IsNotNil(m))
		assert.True(t, zapfluent.IsNotNil(c))
		assert.True(t, zapfluent.IsNotNil(f))
	})
}
