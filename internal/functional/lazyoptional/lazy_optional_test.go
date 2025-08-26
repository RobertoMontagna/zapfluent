package lazyoptional_test

import (
	"strconv"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional/matchers"

	. "github.com/onsi/gomega"
)

func TestLazyOptional_Some(t *testing.T) {
	g := NewWithT(t)

	expectedValue := 42

	opt := lazyoptional.Some(expectedValue)

	g.Expect(opt).To(matchers.BePresent[int]())
	g.Expect(opt).To(matchers.HaveValue(expectedValue))
}

func TestLazyOptional_Empty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Empty[int]()

	g.Expect(opt).To(matchers.BeEmpty[int]())
}

func TestLazyOptional_Filter(t *testing.T) {
	g := NewWithT(t)

	t.Run("on Some with passing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)

		g.Expect(filteredOpt).To(matchers.BePresent[int]())
		g.Expect(filteredOpt).To(matchers.HaveValue(42))
	})

	t.Run("on Some with failing condition", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		predicate := func(i int) bool { return i < 10 }

		filteredOpt := opt.Filter(predicate)

		g.Expect(filteredOpt).To(matchers.BeEmpty[int]())
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		predicate := func(i int) bool { return i > 10 }

		filteredOpt := opt.Filter(predicate)

		g.Expect(filteredOpt).To(matchers.BeEmpty[int]())
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
		f := func(i int) lazyoptional.LazyOptional[string] {
			return lazyoptional.Some(strconv.Itoa(i))
		}

		fmOpt := lazyoptional.FlatMap(opt, f)

		g.Expect(fmOpt).To(matchers.BePresent[string]())
		g.Expect(fmOpt).To(matchers.HaveValue("42"))
	})

	t.Run("on Some that returns Empty", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Empty[string]() }

		fmOpt := lazyoptional.FlatMap(opt, f)

		g.Expect(fmOpt).To(matchers.BeEmpty[string]())
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		f := func(i int) lazyoptional.LazyOptional[string] {
			return lazyoptional.Some(strconv.Itoa(i))
		}

		fmOpt := lazyoptional.FlatMap(opt, f)

		g.Expect(fmOpt).To(matchers.BeEmpty[string]())
	})
}

func TestMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("on Some", func(t *testing.T) {
		opt := lazyoptional.Some(42)
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)

		g.Expect(mappedOpt).To(matchers.BePresent[string]())
		g.Expect(mappedOpt).To(matchers.HaveValue("42"))
	})

	t.Run("on Empty", func(t *testing.T) {
		opt := lazyoptional.Empty[int]()
		mapper := func(i int) string { return strconv.Itoa(i) }

		mappedOpt := lazyoptional.Map(opt, mapper)

		g.Expect(mappedOpt).To(matchers.BeEmpty[string]())
	})
}
