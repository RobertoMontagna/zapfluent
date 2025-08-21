package matchers_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	. "go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional/matchers"
)

func TestBePresent(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		input         any
		shouldSucceed bool
		shouldError   bool
	}{
		{
			name:          "when the lazy optional is present",
			input:         lazyoptional.Some("hello"),
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the lazy optional is empty",
			input:         lazyoptional.Empty[string](),
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not a lazy optional",
			input:         "not-a-lazy-optional",
			shouldSucceed: false,
			shouldError:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			matcher := BePresent[string]()
			success, err := matcher.Match(tc.input)

			if success != tc.shouldSucceed {
				t.Errorf("expected success to be %v, but got %v", tc.shouldSucceed, success)
			}

			if tc.shouldError && err == nil {
				t.Errorf("expected an error, but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
			}
		})
	}
}

func TestBeEmpty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		input         any
		shouldSucceed bool
		shouldError   bool
	}{
		{
			name:          "when the lazy optional is empty",
			input:         lazyoptional.Empty[string](),
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the lazy optional is present",
			input:         lazyoptional.Some("hello"),
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not a lazy optional",
			input:         "not-a-lazy-optional",
			shouldSucceed: false,
			shouldError:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			matcher := BeEmpty[string]()
			success, err := matcher.Match(tc.input)

			if success != tc.shouldSucceed {
				t.Errorf("expected success to be %v, but got %v", tc.shouldSucceed, success)
			}

			if tc.shouldError && err == nil {
				t.Errorf("expected an error, but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
			}
		})
	}
}

func TestHaveValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		input         any
		expectedValue any
		shouldSucceed bool
		shouldError   bool
	}{
		{
			name:          "when the lazy optional has the expected value",
			input:         lazyoptional.Some("hello"),
			expectedValue: "hello",
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the lazy optional has a different value",
			input:         lazyoptional.Some("world"),
			expectedValue: "hello",
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the lazy optional is empty",
			input:         lazyoptional.Empty[string](),
			expectedValue: "hello",
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not a lazy optional",
			input:         "not-a-lazy-optional",
			expectedValue: "hello",
			shouldSucceed: false,
			shouldError:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			matcher := HaveValue[string](tc.expectedValue.(string))
			success, err := matcher.Match(tc.input)

			if success != tc.shouldSucceed {
				t.Errorf("expected success to be %v, but got %v", tc.shouldSucceed, success)
			}

			if tc.shouldError && err == nil {
				t.Errorf("expected an error, but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("did not expect an error, but got: %v", err)
			}
		})
	}
}
