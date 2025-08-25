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

	g.Expect(opt).To(matchers.HaveValue(expectedValue))
}

func TestLazyOptional_Empty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Empty[int]()

	g.Expect(opt).To(matchers.BeEmpty[int]())
}

func TestLazyOptional_Filter_OnSomeWithPassingCondition(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Some(42)
	predicate := func(i int) bool { return i > 10 }

	filteredOpt := opt.Filter(predicate)

	g.Expect(filteredOpt).To(matchers.HaveValue(42))
}

func TestLazyOptional_Filter_OnSomeWithFailingCondition(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Some(42)
	predicate := func(i int) bool { return i < 10 }

	filteredOpt := opt.Filter(predicate)

	g.Expect(filteredOpt).To(matchers.BeEmpty[int]())
}

func TestLazyOptional_Filter_OnEmpty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Empty[int]()
	predicate := func(i int) bool { return i > 10 }

	filteredOpt := opt.Filter(predicate)

	g.Expect(filteredOpt).To(matchers.BeEmpty[int]())
}

func TestFlatMap_OnSomeThatReturnsSome(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Some(42)
	f := func(i int) lazyoptional.LazyOptional[string] {
		return lazyoptional.Some(strconv.Itoa(i))
	}

	fmOpt := lazyoptional.FlatMap(opt, f)

	g.Expect(fmOpt).To(matchers.HaveValue("42"))
}

func TestFlatMap_OnSomeThatReturnsEmpty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Some(42)
	f := func(i int) lazyoptional.LazyOptional[string] { return lazyoptional.Empty[string]() }

	fmOpt := lazyoptional.FlatMap(opt, f)

	g.Expect(fmOpt).To(matchers.BeEmpty[string]())
}

func TestFlatMap_OnEmpty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Empty[int]()
	f := func(i int) lazyoptional.LazyOptional[string] {
		return lazyoptional.Some(strconv.Itoa(i))
	}

	fmOpt := lazyoptional.FlatMap(opt, f)

	g.Expect(fmOpt).To(matchers.BeEmpty[string]())
}

func TestMap_OnSome(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Some(42)
	mapper := func(i int) string { return strconv.Itoa(i) }

	mappedOpt := lazyoptional.Map(opt, mapper)

	g.Expect(mappedOpt).To(matchers.HaveValue("42"))
}

func TestMap_OnEmpty(t *testing.T) {
	g := NewWithT(t)

	opt := lazyoptional.Empty[int]()
	mapper := func(i int) string { return strconv.Itoa(i) }

	mappedOpt := lazyoptional.Map(opt, mapper)

	g.Expect(mappedOpt).To(matchers.BeEmpty[string]())
}
