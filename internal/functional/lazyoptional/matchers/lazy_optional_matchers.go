package matchers

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
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
	// HaveValueFailureMessage is the message returned when the optional does not have the expected value.
	HaveValueFailureMessage = "to have the value"
	// NotHaveValueFailureMessage is the message returned when the optional has the expected value.
	NotHaveValueFailureMessage = "to not have the value"
)

// ErrMatcherWrongType is a sentinel error returned when a matcher receives a value of the wrong type.
var ErrMatcherWrongType = errors.New("matcher received wrong type")

func bePresentWrongTypeError[T any]() error {
	return fmt.Errorf("BePresent matcher expects a lazyoptional.LazyOptional[%T]: %w", *new(T), ErrMatcherWrongType)
}

func beEmptyWrongTypeError[T any]() error {
	return fmt.Errorf("BeEmpty matcher expects a lazyoptional.LazyOptional[%T]: %w", *new(T), ErrMatcherWrongType)
}

func haveValueWrongTypeError[T any]() error {
	return fmt.Errorf("HaveValue matcher expects a lazyoptional.LazyOptional[%T]: %w", *new(T), ErrMatcherWrongType)
}

// BePresent succeeds if the actual value is an optional that is present.
func BePresent[T any]() types.GomegaMatcher {
	return &BePresentMatcher[T]{}
}

type BePresentMatcher[T any] struct{}

func (m *BePresentMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, bePresentWrongTypeError[T]()
	}
	_, isPresent := opt.Get()
	return isPresent, nil
}

func (m *BePresentMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, BePresentFailureMessage)
}

func (m *BePresentMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, NotBePresentFailureMessage)
}

// BeEmpty succeeds if the actual value is an optional that is empty.
func BeEmpty[T any]() types.GomegaMatcher {
	return &BeEmptyMatcher[T]{}
}

type BeEmptyMatcher[T any] struct{}

func (m *BeEmptyMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, beEmptyWrongTypeError[T]()
	}
	_, isPresent := opt.Get()
	return !isPresent, nil
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
		expected: expected,
	}
}

type HaveValueMatcher[T any] struct {
	expected T
}

func (m *HaveValueMatcher[T]) Match(actual any) (bool, error) {
	opt, ok := actual.(lazyoptional.LazyOptional[T])
	if !ok {
		return false, haveValueWrongTypeError[T]()
	}

	val, ok := opt.Get()
	if !ok {
		return false, nil
	}

	return reflect.DeepEqual(val, m.expected), nil
}

func (m *HaveValueMatcher[T]) FailureMessage(actual any) string {
	return format.Message(actual, HaveValueFailureMessage, m.expected)
}

func (m *HaveValueMatcher[T]) NegatedFailureMessage(actual any) string {
	return format.Message(actual, NotHaveValueFailureMessage, m.expected)
}
