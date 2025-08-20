package lazyoptional_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

func TestLazyOptional_Some(t *testing.T) {
	expectedValue := 42

	opt := lazyoptional.Some(expectedValue)
	val, ok := opt.Get()

	assert.True(t, ok)
	assert.Equal(t, expectedValue, val)
}

func TestLazyOptional_Empty(t *testing.T) {
	opt := lazyoptional.Empty[int]()

	_, ok := opt.Get()

	assert.False(t, ok)
}

func TestLazyOptional_Filter(t *testing.T) {
	t.Run("on Some with passing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)
		val, ok := filteredOpt.Get()

		assert.True(t, ok)
		assert.Equal(t, 42, val)
	})

	t.Run("on Some with failing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i < 10 }

		filteredOpt := opt.Filter(predicate)
		_, ok := filteredOpt.Get()

		assert.False(t, ok)
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)
		_, ok := filteredOpt.Get()

		assert.False(t, ok)
	})
}

func TestNewConstantProducer(t *testing.T) {
	expectedV1 := "hello"
	expectedV2 := 42

	producer := lazyoptional.NewConstantProducer(expectedV1, expectedV2)
	v1, v2 := producer()

	assert.Equal(t, expectedV1, v1)
	assert.Equal(t, expectedV2, v2)
}

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
