package matchers

import (
	"errors"
	"fmt"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	optional_matchers "go.robertomontagna.dev/zapfluent/internal/functional/optional/matchers"
)

const (
	// BePresentFailureMessage is the message returned when the optional is not present.
	BePresentFailureMessage = "to be present in the lazy optional"
	// NotBePresentFailureMessage is the message returned when the optional is present.
	NotBePresentFailureMessage = "to not be present in the lazy optional"
	// BeEmptyFailureMessage is the message returned when the optional is not empty.
	BeEmptyFailureMessage = "to be an empty lazy optional"
	// NotBeEmptyFailureMessage is the message returned when the optional is empty.
	NotBeEmptyFailureMessage = "to not be an empty lazy optional"
	// HaveValueFailureMessage is the message returned when the optional does not have the expected
	// value.
	HaveValueFailureMessage = "to have the value"
	// NotHaveValueFailureMessage is the message returned when the optional has the expected value.
	NotHaveValueFailureMessage = "to not have the value"
)

// ErrMatcherWrongType is a sentinel error for when a matcher receives a value of the wrong type.
var ErrMatcherWrongType = errors.New("matcher received wrong type")

// unwrap takes a lazy optional, evaluates it, and wraps the result in an
// eager optional, which can then be used with the existing optional matchers.
// This is the core of the deduplication strategy.
func unwrap[T any](actual any) (optional.Optional[T], error) {
	lazyOpt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return optional.Empty[T](), fmt.Errorf(
			"matcher expects a lazyoptional.LazyOptional[%T], but got %T: %w",
			*new(T),
			actual,
			ErrMatcherWrongType,
		)
	}

	val, isPresent := lazyOpt.Get()
	if !isPresent {
		return optional.Empty[T](), nil
	}
	return optional.Some(val), nil
}

// BePresent succeeds if the actual value is an optional that is present.
func BePresent[T any]() types.GomegaMatcher {
	return &BePresentMatcher[T]{
		delegate: optional_matchers.BePresent[T](),
	}
}

type BePresentMatcher[T any] struct {
	delegate types.GomegaMatcher
}

func (m *BePresentMatcher[T]) Match(actual any) (bool, error) {
	opt, err := unwrap[T](actual)
	if err != nil {
		return false, err
	}
	return m.delegate.Match(opt)
}

func (m *BePresentMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, BePresentFailureMessage)
}

func (m *BePresentMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, NotBePresentFailureMessage)
}

// BeEmpty succeeds if the actual value is an optional that is empty.
func BeEmpty[T any]() types.GomegaMatcher {
	return &BeEmptyMatcher[T]{
		delegate: optional_matchers.BeEmpty[T](),
	}
}

type BeEmptyMatcher[T any] struct {
	delegate types.GomegaMatcher
}

func (m *BeEmptyMatcher[T]) Match(actual any) (bool, error) {
	opt, err := unwrap[T](actual)
	if err != nil {
		return false, err
	}
	return m.delegate.Match(opt)
}

func (m *BeEmptyMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, BeEmptyFailureMessage)
}

func (m *BeEmptyMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, NotBeEmptyFailureMessage)
}

// HaveValue succeeds if the actual value is an optional that is present and
// contains the expected value.
func HaveValue[T any](expected T) types.GomegaMatcher {
	return &HaveValueMatcher[T]{
		delegate: optional_matchers.HaveValue(expected),
	}
}

type HaveValueMatcher[T any] struct {
	delegate types.GomegaMatcher
}

func (m *HaveValueMatcher[T]) Match(actual any) (bool, error) {
	opt, err := unwrap[T](actual)
	if err != nil {
		return false, err
	}
	return m.delegate.Match(opt)
}

func (m *HaveValueMatcher[T]) FailureMessage(actual any) string {
	delegateWithValue := m.delegate.(*optional_matchers.HaveValueMatcher[T])
	return format.Message(actual, HaveValueFailureMessage, delegateWithValue.Expected())
}

func (m *HaveValueMatcher[T]) NegatedFailureMessage(actual any) string {
	delegateWithValue := m.delegate.(*optional_matchers.HaveValueMatcher[T])
	return format.Message(actual, NotHaveValueFailureMessage, delegateWithValue.Expected())
}
