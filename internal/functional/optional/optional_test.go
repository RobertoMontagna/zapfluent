package optional_test

import (
	"errors"
	"strconv"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/contracts/matchers"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional"

	. "github.com/onsi/gomega"
)

func TestOptional_Some(t *testing.T) {
	g := NewWithT(t)

	o := optional.Some("test")

	g.Expect(o).To(matchers.BePresent[string]())
	g.Expect(o).To(matchers.HaveValue("test"))
}

func TestOptional_Map_WhenMapperReturnsNil(t *testing.T) {
	g := NewWithT(t)

	opt := optional.OfError(errors.New("test error"))

	res := optional.Map(opt, func(t error) error {
		return nil
	})

	g.Expect(res).To(matchers.BeEmpty[error]())
}

func TestOptional_Empty(t *testing.T) {
	g := NewWithT(t)

	o := optional.Empty[string]()

	g.Expect(o).To(matchers.BeEmpty[string]())
}

func TestOptional_OfPtr(t *testing.T) {
	val := 123
	testCases := []struct {
		name          string
		pointer       *int
		shouldBeEmpty bool
		expectedValue int
	}{
		{
			name:          "with nil pointer",
			pointer:       nil,
			shouldBeEmpty: true,
		},
		{
			name:          "with non-nil pointer",
			pointer:       &val,
			shouldBeEmpty: false,
			expectedValue: 123,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			o := optional.OfPtr(tc.pointer)

			if tc.shouldBeEmpty {
				g.Expect(o).To(matchers.BeEmpty[int]())
			} else {
				g.Expect(o).To(matchers.BePresent[int]())
				g.Expect(o).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestOptional_OfError(t *testing.T) {
	testErr := errors.New("test error")
	testCases := []struct {
		name          string
		err           error
		shouldBeEmpty bool
		expectedValue error
	}{
		{
			name:          "with nil error",
			err:           nil,
			shouldBeEmpty: true,
		},
		{
			name:          "with non-nil error",
			err:           testErr,
			shouldBeEmpty: false,
			expectedValue: testErr,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			o := optional.OfError(tc.err)

			if tc.shouldBeEmpty {
				g.Expect(o).To(matchers.BeEmpty[error]())
			} else {
				g.Expect(o).To(matchers.BePresent[error]())
				g.Expect(o).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestOptional_Map(t *testing.T) {
	testCases := []struct {
		name          string
		initialValue  optional.Optional[int]
		mapper        func(int) string
		shouldBeEmpty bool
		expectedValue string
	}{
		{
			name:          "with present value",
			initialValue:  optional.Some(123),
			mapper:        strconv.Itoa,
			shouldBeEmpty: false,
			expectedValue: "123",
		},
		{
			name:          "with empty value",
			initialValue:  optional.Empty[int](),
			mapper:        strconv.Itoa,
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			mapped := optional.Map(tc.initialValue, tc.mapper)

			if tc.shouldBeEmpty {
				g.Expect(mapped).To(matchers.BeEmpty[string]())
			} else {
				g.Expect(mapped).To(matchers.BePresent[string]())
				g.Expect(mapped).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestOptional_Get(t *testing.T) {
	testCases := []struct {
		name          string
		initialValue  optional.Optional[string]
		expectedOK    bool
		expectedValue string
	}{
		{
			name:          "with present value",
			initialValue:  optional.Some("hello"),
			expectedOK:    true,
			expectedValue: "hello",
		},
		{
			name:         "with empty value",
			initialValue: optional.Empty[string](),
			expectedOK:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			val, ok := tc.initialValue.Get()

			g.Expect(ok).To(Equal(tc.expectedOK))
			if tc.expectedOK {
				g.Expect(val).To(Equal(tc.expectedValue))
			} else {
				g.Expect(val).To(BeZero())
			}
		})
	}
}

func TestOptional_FlatMap(t *testing.T) {
	testCases := []struct {
		name          string
		initialValue  optional.Optional[int]
		f             func(int) optional.Optional[string]
		shouldBeEmpty bool
		expectedValue string
	}{
		{
			name:         "with present value mapping to present",
			initialValue: optional.Some(123),
			f: func(i int) optional.Optional[string] {
				return optional.Some(strconv.Itoa(i))
			},
			shouldBeEmpty: false,
			expectedValue: "123",
		},
		{
			name:         "with present value mapping to empty",
			initialValue: optional.Some(123),
			f: func(i int) optional.Optional[string] {
				return optional.Empty[string]()
			},
			shouldBeEmpty: true,
		},
		{
			name:         "with empty value",
			initialValue: optional.Empty[int](),
			f: func(i int) optional.Optional[string] {
				return optional.Some(strconv.Itoa(i))
			},
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			mapped := optional.FlatMap(tc.initialValue, tc.f)

			if tc.shouldBeEmpty {
				g.Expect(mapped).To(matchers.BeEmpty[string]())
			} else {
				g.Expect(mapped).To(matchers.BePresent[string]())
				g.Expect(mapped).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}
