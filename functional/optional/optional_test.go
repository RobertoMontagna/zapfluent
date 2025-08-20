package optional_test

import (
	"strconv"
	"testing"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/functional/optional"
)

func TestOptional_Some(t *testing.T) {
	g := NewWithT(t)
	o := optional.Some("test")
	g.Expect(o.IsPresent()).To(BeTrue())
	val, ok := o.Get()
	g.Expect(ok).To(BeTrue())
	g.Expect(val).To(Equal("test"))
}

func TestOptional_Empty(t *testing.T) {
	g := NewWithT(t)
	o := optional.Empty[string]()
	g.Expect(o.IsPresent()).To(BeFalse())
	val, ok := o.Get()
	g.Expect(ok).To(BeFalse())
	g.Expect(val).To(Equal("")) // Zero value
}

func TestOptional_ForEach(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value", func(t *testing.T) {
		o := optional.Some("test")
		var result string
		o.ForEach(func(s string) {
			result = s
		})
		g.Expect(result).To(Equal("test"))
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[string]()
		var result string
		o.ForEach(func(s string) {
			result = "should not be called"
		})
		g.Expect(result).To(BeEmpty())
	})
}

func TestOptional_Map(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.Map(o, strconv.Itoa)
		g.Expect(mapped.IsPresent()).To(BeTrue())
		val, _ := mapped.Get()
		g.Expect(val).To(Equal("123"))
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.Map(o, strconv.Itoa)
		g.Expect(mapped.IsPresent()).To(BeFalse())
	})
}

func TestOptional_FlatMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value mapping to present", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		g.Expect(mapped.IsPresent()).To(BeTrue())
		val, _ := mapped.Get()
		g.Expect(val).To(Equal("123"))
	})

	t.Run("with present value mapping to empty", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Empty[string]()
		})
		g.Expect(mapped.IsPresent()).To(BeFalse())
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		g.Expect(mapped.IsPresent()).To(BeFalse())
	})
}
