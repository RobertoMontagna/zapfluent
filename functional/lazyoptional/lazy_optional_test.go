package lazyoptional_test

import (
	"strconv"
	"testing"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

func TestLazyOptional_Some(t *testing.T) {
	g := NewWithT(t)
	expectedValue := 42

	opt := lazyoptional.Some(expectedValue)
	val, ok := opt.Get()

	g.Expect(ok).To(BeTrue())
	g.Expect(val).To(Equal(expectedValue))
}

func TestLazyOptional_Empty(t *testing.T) {
	g := NewWithT(t)
	opt := lazyoptional.Empty[int]()

	_, ok := opt.Get()

	g.Expect(ok).To(BeFalse())
}

func TestLazyOptional_Filter(t *testing.T) {
	g := NewWithT(t)

	t.Run("on Some with passing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)
		val, ok := filteredOpt.Get()

		g.Expect(ok).To(BeTrue())
		g.Expect(val).To(Equal(42))
	})

	t.Run("on Some with failing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i < 10 }

		filteredOpt := opt.Filter(predicate)
		_, ok := filteredOpt.Get()

		g.Expect(ok).To(BeFalse())
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)
		_, ok := filteredOpt.Get()

		g.Expect(ok).To(BeFalse())
	})
}

func TestNewConstantProducer(t *testing.T) {
	g := NewWithT(t)
	expectedV1 := "hello"
	expectedV2 := 42

	producer := lazyoptional.NewConstantProducer(expectedV1, expectedV2)
	v1, v2 := producer()

	g.Expect(v1).To(Equal(expectedV1))
	g.Expect(v2).To(Equal(expectedV2))
}

func TestFlatMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("on Some that returns Some", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Some(strconv.Itoa(i)) }

		fmOpt := lazyoptional.FlatMap(opt, f)
		val, ok := fmOpt.Get()

		g.Expect(ok).To(BeTrue())
		g.Expect(val).To(Equal("42"))
	})

	t.Run("on Some that returns Empty", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Empty[string]() }

		fmOpt := lazyoptional.FlatMap(opt, f)
		_, ok := fmOpt.Get()

		g.Expect(ok).To(BeFalse())
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Some(strconv.Itoa(i)) }

		fmOpt := lazyoptional.FlatMap(opt, f)
		_, ok := fmOpt.Get()

		g.Expect(ok).To(BeFalse())
	})
}

func TestMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("on Some", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)
		val, ok := mappedOpt.Get()

		g.Expect(ok).To(BeTrue())
		g.Expect(val).To(Equal("42"))
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)
		_, ok := mappedOpt.Get()

		g.Expect(ok).To(BeFalse())
	})
}
