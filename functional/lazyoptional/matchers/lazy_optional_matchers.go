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

func (m *bePresentMatcher) Match(actual interface{}) (bool, error) {
	val := reflect.ValueOf(actual)
	getMethod := val.MethodByName("Get")
	if !getMethod.IsValid() {
		return false, fmt.Errorf("BePresent matcher expects a type with a Get() method, but got %T", actual)
	}

	results := getMethod.Call(nil)
	if len(results) != 2 || results[1].Kind() != reflect.Bool {
		return false, fmt.Errorf("Get() method must return a value and a boolean")
	}

	return results[1].Bool(), nil
}

func (m *bePresentMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be present")
}

func (m *bePresentMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be present")
}

// BeEmpty succeeds if the actual value is an optional that is empty.
func BeEmpty() types.GomegaMatcher {
	return &beEmptyMatcher{}
}

type beEmptyMatcher struct{}

func (m *beEmptyMatcher) Match(actual interface{}) (bool, error) {
	val := reflect.ValueOf(actual)
	getMethod := val.MethodByName("Get")
	if !getMethod.IsValid() {
		return false, fmt.Errorf("BeEmpty matcher expects a type with a Get() method, but got %T", actual)
	}

	results := getMethod.Call(nil)
	if len(results) != 2 || results[1].Kind() != reflect.Bool {
		return false, fmt.Errorf("Get() method must return a value and a boolean")
	}

	return !results[1].Bool(), nil
}

func (m *beEmptyMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to be empty")
}

func (m *beEmptyMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to be empty")
}

// HaveValue succeeds if the actual value is an optional that is present and
// contains the expected value.
func HaveValue(expected interface{}) types.GomegaMatcher {
	return &haveValueMatcher{
		expected: expected,
	}
}

type haveValueMatcher struct {
	expected interface{}
}

func (m *haveValueMatcher) Match(actual interface{}) (bool, error) {
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

func (m *haveValueMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, "to have value", m.expected)
}

func (m *haveValueMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, "not to have value", m.expected)
}
