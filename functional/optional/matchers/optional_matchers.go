package matchers

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// BePresent succeeds if the actual value is an optional that is present.
func BePresent() types.GomegaMatcher {
	return &bePresentMatcher{}
}

type bePresentMatcher struct{}

func (m *bePresentMatcher) Match(actual any) (bool, error) {
	val := reflect.ValueOf(actual)
	isPresentMethod := val.MethodByName("IsPresent")
	if !isPresentMethod.IsValid() {
		return false, fmt.Errorf("BePresent matcher expects a type with an IsPresent() method, but got %T", actual)
	}

	results := isPresentMethod.Call(nil)
	if len(results) != 1 || results[0].Kind() != reflect.Bool {
		return false, fmt.Errorf("IsPresent() method must return a single boolean value")
	}

	return results[0].Bool(), nil
}

func (m *bePresentMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to be present")
}

func (m *bePresentMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to be present")
}

// BeEmpty succeeds if the actual value is an optional that is empty.
func BeEmpty() types.GomegaMatcher {
	return &beEmptyMatcher{}
}

type beEmptyMatcher struct{}

func (m *beEmptyMatcher) Match(actual any) (bool, error) {
	val := reflect.ValueOf(actual)
	isPresentMethod := val.MethodByName("IsPresent")
	if !isPresentMethod.IsValid() {
		return false, fmt.Errorf("BeEmpty matcher expects a type with an IsPresent() method, but got %T", actual)
	}

	results := isPresentMethod.Call(nil)
	if len(results) != 1 || results[0].Kind() != reflect.Bool {
		return false, fmt.Errorf("IsPresent() method must return a single boolean value")
	}

	return !results[0].Bool(), nil
}

func (m *beEmptyMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to be empty")
}

func (m *beEmptyMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to be empty")
}

// HaveValue succeeds if the actual value is an optional that is present and
// contains the expected value.
func HaveValue(expected any) types.GomegaMatcher {
	return &haveValueMatcher{
		expected: expected,
	}
}

type haveValueMatcher struct {
	expected any
}

func (m *haveValueMatcher) Match(actual any) (bool, error) {
	val := reflect.ValueOf(actual)
	getMethod := val.MethodByName("Get")
	if !getMethod.IsValid() {
		return false, fmt.Errorf("HaveValue matcher expects a type with a Get() method, but got %T", actual)
	}

	results := getMethod.Call(nil)
	if len(results) != 2 || results[1].Kind() != reflect.Bool {
		return false, fmt.Errorf("Get() method must return a value and a boolean")
	}

	// Check if the optional is present
	if !results[1].Bool() {
		return false, nil
	}

	// Check if the value matches the expected value
	return reflect.DeepEqual(results[0].Interface(), m.expected), nil
}

func (m *haveValueMatcher) FailureMessage(actual any) string {
	return format.Message(actual, "to have value", m.expected)
}

func (m *haveValueMatcher) NegatedFailureMessage(actual any) string {
	return format.Message(actual, "not to have value", m.expected)
}
