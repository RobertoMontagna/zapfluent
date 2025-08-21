package optional_test

import (
	"errors"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional/matchers"

	. "github.com/onsi/gomega"
)

func TestLiftToOptional(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func() *int { return nil }
		lifted := optional.LiftToOptional(f)

		// Act
		result := lifted()

		// Assert
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		// Arrange
		v := 123
		f := func() *int { return &v }
		lifted := optional.LiftToOptional(f)

		// Act
		result := lifted()

		// Assert
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftToOptional1(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string) *int { return nil }
		lifted := optional.LiftToOptional1(f)

		// Act
		result := lifted("test")

		// Assert
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		// Arrange
		v := 123
		f := func(s string) *int {
			g.Expect(s).To(Equal("test"))
			return &v
		}
		lifted := optional.LiftToOptional1(f)

		// Act
		result := lifted("test")

		// Assert
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftToOptional2(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string, i int) *int { return nil }
		lifted := optional.LiftToOptional2(f)

		// Act
		result := lifted("test", 1)

		// Assert
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		// Arrange
		v := 123
		f := func(s string, i int) *int {
			g.Expect(s).To(Equal("test"))
			g.Expect(i).To(Equal(1))
			return &v
		}
		lifted := optional.LiftToOptional2(f)

		// Act
		result := lifted("test", 1)

		// Assert
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftErrorToOptional(t *testing.T) {
	g := NewWithT(t)

	// Arrange
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func() error { return nil }
		lifted := optional.LiftErrorToOptional(f)

		// Act
		result := lifted()

		// Assert
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		// Arrange
		f := func() error { return testErr }
		lifted := optional.LiftErrorToOptional(f)

		// Act
		result := lifted()

		// Assert
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}

func TestLiftErrorToOptional1(t *testing.T) {
	g := NewWithT(t)

	// Arrange
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string) error { return nil }
		lifted := optional.LiftErrorToOptional1(f)

		// Act
		result := lifted("test")

		// Assert
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string) error {
			g.Expect(s).To(Equal("test"))
			return testErr
		}
		lifted := optional.LiftErrorToOptional1(f)

		// Act
		result := lifted("test")

		// Assert
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}

func TestLiftErrorToOptional2(t *testing.T) {
	g := NewWithT(t)

	// Arrange
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string, i int) error { return nil }
		lifted := optional.LiftErrorToOptional2(f)

		// Act
		result := lifted("test", 1)

		// Assert
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		// Arrange
		f := func(s string, i int) error {
			g.Expect(s).To(Equal("test"))
			g.Expect(i).To(Equal(1))
			return testErr
		}
		lifted := optional.LiftErrorToOptional2(f)

		// Act
		result := lifted("test", 1)

		// Assert
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}
