package optional_test

import (
	"errors"
	"strconv"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional/matchers"

	. "github.com/onsi/gomega"
)

func TestOptional_Some(t *testing.T) {
	g := NewWithT(t)

	// Act
	o := optional.Some("test")

	// Assert
	g.Expect(o).To(matchers.BePresent[string]())
	g.Expect(o).To(matchers.HaveValue("test"))
}

func TestOptional_Empty(t *testing.T) {
	g := NewWithT(t)

	// Act
	o := optional.Empty[string]()

	// Assert
	g.Expect(o).To(matchers.BeEmpty[string]())
}

func TestOptional_OfPtr(t *testing.T) {
	g := NewWithT(t)

	t.Run("with nil pointer", func(t *testing.T) {
		// Act
		o := optional.OfPtr[int](nil)

		// Assert
		g.Expect(o).To(matchers.BeEmpty[int]())
	})

	t.Run("with non-nil pointer", func(t *testing.T) {
		// Arrange
		v := 123

		// Act
		o := optional.OfPtr(&v)

		// Assert
		g.Expect(o).To(matchers.BePresent[int]())
		g.Expect(o).To(matchers.HaveValue(123))
	})
}

func TestOptional_OfError(t *testing.T) {
	g := NewWithT(t)

	// Arrange
	testErr := errors.New("test error")

	t.Run("with nil error", func(t *testing.T) {
		// Act
		o := optional.OfError(nil)

		// Assert
		g.Expect(o).To(matchers.BeEmpty[error]())
	})

	t.Run("with non-nil error", func(t *testing.T) {
		// Act
		o := optional.OfError(testErr)

		// Assert
		g.Expect(o).To(matchers.BePresent[error]())
		g.Expect(o).To(matchers.HaveValue(testErr))
	})
}

func TestOptional_Map(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value", func(t *testing.T) {
		// Arrange
		o := optional.Some(123)

		// Act
		mapped := optional.Map(o, strconv.Itoa)

		// Assert
		g.Expect(mapped).To(matchers.BePresent[string]())
		g.Expect(mapped).To(matchers.HaveValue("123"))
	})

	t.Run("with empty value", func(t *testing.T) {
		// Arrange
		o := optional.Empty[int]()

		// Act
		mapped := optional.Map(o, strconv.Itoa)

		// Assert
		g.Expect(mapped).To(matchers.BeEmpty[string]())
	})
}

func TestOptional_FlatMap(t *testing.T) {
	g := NewWithT(t)

	t.Run("with present value mapping to present", func(t *testing.T) {
		// Arrange
		o := optional.Some(123)

		// Act
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})

		// Assert
		g.Expect(mapped).To(matchers.BePresent[string]())
		g.Expect(mapped).To(matchers.HaveValue("123"))
	})

	t.Run("with present value mapping to empty", func(t *testing.T) {
		// Arrange
		o := optional.Some(123)

		// Act
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Empty[string]()
		})

		// Assert
		g.Expect(mapped).To(matchers.BeEmpty[string]())
	})

	t.Run("with empty value", func(t *testing.T) {
		// Arrange
		o := optional.Empty[int]()

		// Act
		mapped := optional.FlatMap(o, func(i int) optional.Optional[string] {
			return optional.Some(strconv.Itoa(i))
		})

		// Assert
		g.Expect(mapped).To(matchers.BeEmpty[string]())
	})
}
