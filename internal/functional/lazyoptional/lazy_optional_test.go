package lazyoptional_test

import (
	"strconv"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/contracts/matchers"
	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"

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
	testCases := []struct {
		name          string
		initialValue  lazyoptional.LazyOptional[int]
		predicate     func(int) bool
		shouldBeEmpty bool
		expectedValue int
	}{
		{
			name:          "on Some with passing condition",
			initialValue:  lazyoptional.Some(42),
			predicate:     func(i int) bool { return i > 10 },
			shouldBeEmpty: false,
			expectedValue: 42,
		},
		{
			name:          "on Some with failing condition",
			initialValue:  lazyoptional.Some(42),
			predicate:     func(i int) bool { return i < 10 },
			shouldBeEmpty: true,
		},
		{
			name:          "on Empty",
			initialValue:  lazyoptional.Empty[int](),
			predicate:     func(i int) bool { return i > 10 },
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			filteredOpt := tc.initialValue.Filter(tc.predicate)

			if tc.shouldBeEmpty {
				g.Expect(filteredOpt).To(matchers.BeEmpty[int]())
			} else {
				g.Expect(filteredOpt).To(matchers.BePresent[int]())
				g.Expect(filteredOpt).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
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
	testCases := []struct {
		name          string
		initialValue  lazyoptional.LazyOptional[int]
		f             func(int) lazyoptional.LazyOptional[string]
		shouldBeEmpty bool
		expectedValue string
	}{
		{
			name:         "on Some that returns Some",
			initialValue: lazyoptional.Some(42),
			f: func(i int) lazyoptional.LazyOptional[string] {
				return lazyoptional.Some(strconv.Itoa(i))
			},
			shouldBeEmpty: false,
			expectedValue: "42",
		},
		{
			name:         "on Some that returns Empty",
			initialValue: lazyoptional.Some(42),
			f: func(i int) lazyoptional.LazyOptional[string] {
				return lazyoptional.Empty[string]()
			},
			shouldBeEmpty: true,
		},
		{
			name:         "on Empty",
			initialValue: lazyoptional.Empty[int](),
			f: func(i int) lazyoptional.LazyOptional[string] {
				return lazyoptional.Some(strconv.Itoa(i))
			},
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			fmOpt := lazyoptional.FlatMap(tc.initialValue, tc.f)

			if tc.shouldBeEmpty {
				g.Expect(fmOpt).To(matchers.BeEmpty[string]())
			} else {
				g.Expect(fmOpt).To(matchers.BePresent[string]())
				g.Expect(fmOpt).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name          string
		initialValue  lazyoptional.LazyOptional[int]
		mapper        func(int) string
		shouldBeEmpty bool
		expectedValue string
	}{
		{
			name:          "on Some",
			initialValue:  lazyoptional.Some(42),
			mapper:        func(i int) string { return strconv.Itoa(i) },
			shouldBeEmpty: false,
			expectedValue: "42",
		},
		{
			name:          "on Empty",
			initialValue:  lazyoptional.Empty[int](),
			mapper:        func(i int) string { return strconv.Itoa(i) },
			shouldBeEmpty: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			mappedOpt := lazyoptional.Map(tc.initialValue, tc.mapper)

			if tc.shouldBeEmpty {
				g.Expect(mappedOpt).To(matchers.BeEmpty[string]())
			} else {
				g.Expect(mappedOpt).To(matchers.BePresent[string]())
				g.Expect(mappedOpt).To(matchers.HaveValue(tc.expectedValue))
			}
		})
	}
}
