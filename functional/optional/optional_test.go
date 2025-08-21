package optional_test

import (
	"errors"
	"strconv"
	"testing"

	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/functional/optional"
	"go.robertomontagna.dev/zapfluent/functional/optional/matchers"
)

func TestOptional_Some(t *testing.T) {
	g := NewWithT(t)
	o := optional.Some("test")
	g.Expect(o).To(matchers.BePresent())
	g.Expect(o).To(matchers.HaveValue("test"))
}

func TestOptional_Empty(t *testing.T) {
	g := NewWithT(t)
	o := optional.Empty[string]()
	g.Expect(o).To(matchers.BeEmpty())
}

func TestOptional_OfPtr(t *testing.T) {
	g := NewWithT(t)

	t.Run("with nil pointer", func(t *testing.T) {
		o := optional.OfPtr[int](nil)
		g.Expect(o).To(matchers.BeEmpty())
	})

	t.Run("with non-nil pointer", func(t *testing.T) {
		v := 123
		o := optional.OfPtr(&v)
		g.Expect(o).To(matchers.BePresent())
		g.Expect(o).To(matchers.HaveValue(123))
	})
}

func TestOptional_OfError(t *testing.T) {
	g := NewWithT(t)
	testErr := errors.New("test error")

	t.Run("with nil error", func(t *testing.T) {
		o := optional.OfError(nil)
		g.Expect(o).To(matchers.BeEmpty())
	})

	t.Run("with non-nil error", func(t *testing.T) {
		o := optional.OfError(testErr)
		g.Expect(o).To(matchers.BePresent())
		g.Expect(o).To(matchers.HaveValue(testErr))
	})
}

func TestOptional_Map(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.Map(o, strconv.Itoa)
		g.Expect(mapped).To(matchers.BePresent())
		g.Expect(mapped).To(matchers.HaveValue("123"))
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.Map(o, strconv.Itoa)
		g.Expect(mapped).To(matchers.BeEmpty())
	})
}

func TestOptional_FlatMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value mapping to present", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		g.Expect(mapped).To(matchers.BePresent())
		g.Expect(mapped).To(matchers.HaveValue("123"))
	})

	t.Run("with present value mapping to empty", func(t *testing.T) {
		o := optional.Some(123)
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Empty[string]()
		})
		g.Expect(mapped).To(matchers.BeEmpty())
	})

	t.Run("with empty value", func(t *testing.T) {
		o := optional.Empty[int]()
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})
		g.Expect(mapped).To(matchers.BeEmpty())
	})
}
