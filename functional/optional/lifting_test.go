package optional_test

import (
	"errors"
	"testing"

	"go.robertomontagna.dev/zapfluent/functional/optional"
	"go.robertomontagna.dev/zapfluent/functional/optional/matchers"

	. "github.com/onsi/gomega"
)

func TestLiftToOptional(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func() *int { return nil }
		lifted := optional.LiftToOptional(f)
		result := lifted()
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		v := 123
		f := func() *int { return &v }
		lifted := optional.LiftToOptional(f)
		result := lifted()
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftToOptional1(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func(s string) *int { return nil }
		lifted := optional.LiftToOptional1(f)
		result := lifted("test")
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		v := 123
		f := func(s string) *int {
			g.Expect(s).To(Equal("test"))
			return &v
		}
		lifted := optional.LiftToOptional1(f)
		result := lifted("test")
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftToOptional2(t *testing.T) {
	g := NewWithT(t)

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func(s string, i int) *int { return nil }
		lifted := optional.LiftToOptional2(f)
		result := lifted("test", 1)
		g.Expect(result).To(matchers.BeEmpty[int]())
	})

	t.Run("for value-returning function", func(t *testing.T) {
		v := 123
		f := func(s string, i int) *int {
			g.Expect(s).To(Equal("test"))
			g.Expect(i).To(Equal(1))
			return &v
		}
		lifted := optional.LiftToOptional2(f)
		result := lifted("test", 1)
		g.Expect(result).To(matchers.BePresent[int]())
		g.Expect(result).To(matchers.HaveValue(123))
	})
}

func TestLiftErrorToOptional(t *testing.T) {
	g := NewWithT(t)
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func() error { return nil }
		lifted := optional.LiftErrorToOptional(f)
		result := lifted()
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		f := func() error { return testErr }
		lifted := optional.LiftErrorToOptional(f)
		result := lifted()
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}

func TestLiftErrorToOptional1(t *testing.T) {
	g := NewWithT(t)
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func(s string) error { return nil }
		lifted := optional.LiftErrorToOptional1(f)
		result := lifted("test")
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		f := func(s string) error {
			g.Expect(s).To(Equal("test"))
			return testErr
		}
		lifted := optional.LiftErrorToOptional1(f)
		result := lifted("test")
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}

func TestLiftErrorToOptional2(t *testing.T) {
	g := NewWithT(t)
	testErr := errors.New("test error")

	t.Run("for nil-returning function", func(t *testing.T) {
		f := func(s string, i int) error { return nil }
		lifted := optional.LiftErrorToOptional2(f)
		result := lifted("test", 1)
		g.Expect(result).To(matchers.BeEmpty[error]())
	})

	t.Run("for error-returning function", func(t *testing.T) {
		f := func(s string, i int) error {
			g.Expect(s).To(Equal("test"))
			g.Expect(i).To(Equal(1))
			return testErr
		}
		lifted := optional.LiftErrorToOptional2(f)
		result := lifted("test", 1)
		g.Expect(result).To(matchers.BePresent[error]())
		g.Expect(result).To(matchers.HaveValue(testErr))
	})
}
