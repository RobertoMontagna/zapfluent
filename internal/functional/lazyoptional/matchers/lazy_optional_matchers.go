package matchers

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
)

// BePresent succeeds if the actual value is an optional that is present.
func BePresent[T any]() types.GomegaMatcher {
	return &BePresentMatcher[T]{}
}

type BePresentMatcher[T any] struct{}

func (m *BePresentMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, fmt.Errorf("BePresent matcher expects an lazyoptional.LazyOptional[%T]", *new(T))
	}
	_, isPresent := opt.Get()
	return isPresent, nil
}

func (m *BePresentMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, "to be present")
}

func (m *BePresentMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to be present")
}

// BeEmpty succeeds if the actual value is an optional that is empty.
func BeEmpty[T any]() types.GomegaMatcher {
	return &BeEmptyMatcher[T]{}
}

type BeEmptyMatcher[T any] struct{}

func (m *BeEmptyMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, fmt.Errorf("BeEmpty matcher expects an lazyoptional.LazyOptional[%T]", *new(T))
	}
	_, isPresent := opt.Get()
	return !isPresent, nil
}

func (m *BeEmptyMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, "to be empty")
}

func (m *BeEmptyMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to be empty")
}

// HaveValue succeeds if the actual value is an optional that is present and
// contains the expected value.
func HaveValue[T any](expected T) types.GomegaMatcher {
	return &HaveValueMatcher[T]{
		expected: expected,
	}
}

type HaveValueMatcher[T any] struct {
	expected T
}

func (m *HaveValueMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, fmt.Errorf("HaveValue matcher expects an lazyoptional.LazyOptional[%T]", *new(T))
	}

	val, ok := opt.Get()
	if !ok {
		return false, nil
	}

	return reflect.DeepEqual(val, m.expected), nil
}

func (m *HaveValueMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, "to have value", m.expected)
}

func (m *HaveValueMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to have value", m.expected)
}
