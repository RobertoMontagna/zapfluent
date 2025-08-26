package optional_test

import (
	"errors"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/contracts/matchers"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional"

	. "github.com/onsi/gomega"
)

func TestLiftToOptional(t *testing.T) {
	v := 123
	testCases := []struct {
		name          string
		f             func() *int
		shouldBeEmpty bool
		expectedValue int
	}{
		{
			name:          "for nil-returning function",
			f:             func() *int { return nil },
			shouldBeEmpty: true,
		},
		{
			name:          "for value-returning function",
			f:             func() *int { return &v },
			shouldBeEmpty: false,
			expectedValue: 123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftToOptional(tc.f)

			result := lifted()

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[int]())
			} else {
				g.Expect(result).To(matchers.BePresent[int]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestLiftToOptional1(t *testing.T) {
	v := 123
	testCases := []struct {
		name          string
		f             func(string) *int
		shouldBeEmpty bool
		expectedValue int
	}{
		{
			name:          "for nil-returning function",
			f:             func(s string) *int { return nil },
			shouldBeEmpty: true,
		},
		{
			name: "for value-returning function",
			f: func(s string) *int {
				return &v
			},
			shouldBeEmpty: false,
			expectedValue: 123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftToOptional1(tc.f)

			result := lifted("test")

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[int]())
			} else {
				g.Expect(result).To(matchers.BePresent[int]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestLiftToOptional2(t *testing.T) {
	v := 123
	testCases := []struct {
		name          string
		f             func(string, int) *int
		shouldBeEmpty bool
		expectedValue int
	}{
		{
			name:          "for nil-returning function",
			f:             func(s string, i int) *int { return nil },
			shouldBeEmpty: true,
		},
		{
			name: "for value-returning function",
			f: func(s string, i int) *int {
				return &v
			},
			shouldBeEmpty: false,
			expectedValue: 123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftToOptional2(tc.f)

			result := lifted("test", 1)

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[int]())
			} else {
				g.Expect(result).To(matchers.BePresent[int]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestLiftErrorToOptional(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name          string
		f             func() error
		shouldBeEmpty bool
		expectedValue error
	}{
		{
			name:          "for nil-returning function",
			f:             func() error { return nil },
			shouldBeEmpty: true,
		},
		{
			name:          "for error-returning function",
			f:             func() error { return testErr },
			shouldBeEmpty: false,
			expectedValue: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftErrorToOptional(tc.f)

			result := lifted()

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[error]())
			} else {
				g.Expect(result).To(matchers.BePresent[error]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestLiftErrorToOptional1(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name          string
		f             func(string) error
		shouldBeEmpty bool
		expectedValue error
	}{
		{
			name:          "for nil-returning function",
			f:             func(s string) error { return nil },
			shouldBeEmpty: true,
		},
		{
			name: "for error-returning function",
			f: func(s string) error {
				return testErr
			},
			shouldBeEmpty: false,
			expectedValue: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftErrorToOptional1(tc.f)

			result := lifted("test")

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[error]())
			} else {
				g.Expect(result).To(matchers.BePresent[error]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestLiftErrorToOptional2(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name          string
		f             func(string, int) error
		shouldBeEmpty bool
		expectedValue error
	}{
		{
			name:          "for nil-returning function",
			f:             func(s string, i int) error { return nil },
			shouldBeEmpty: true,
		},
		{
			name: "for error-returning function",
			f: func(s string, i int) error {
				return testErr
			},
			shouldBeEmpty: false,
			expectedValue: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			lifted := optional.LiftErrorToOptional2(tc.f)

			result := lifted("test", 1)

			if tc.shouldBeEmpty {
				g.Expect(result).To(matchers.BeEmpty[error]())
			} else {
				g.Expect(result).To(matchers.BePresent[error]())
				g.Expect(result).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}
