package matchers_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"

	. "go.robertomontagna.dev/zapfluent/internal/functional/optional/matchers"
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
			name:          "when the optional is present",
			input:         optional.Some("hello"),
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the optional is empty",
			input:         optional.Empty[string](),
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not an optional",
			input:         "not-an-optional",
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
			name:          "when the optional is empty",
			input:         optional.Empty[string](),
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the optional is present",
			input:         optional.Some("hello"),
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not an optional",
			input:         "not-an-optional",
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
			name:          "when the optional has the expected value",
			input:         optional.Some("hello"),
			expectedValue: "hello",
			shouldSucceed: true,
			shouldError:   false,
		},
		{
			name:          "when the optional has a different value",
			input:         optional.Some("world"),
			expectedValue: "hello",
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the optional is empty",
			input:         optional.Empty[string](),
			expectedValue: "hello",
			shouldSucceed: false,
			shouldError:   false,
		},
		{
			name:          "when the actual is not an optional",
			input:         "not-an-optional",
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
