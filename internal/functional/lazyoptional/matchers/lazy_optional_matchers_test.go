package matchers_test

import (
	"strings"
	"testing"

	"github.com/onsi/gomega/types"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional/matchers"
	"go.robertomontagna.dev/zapfluent/testutil"

	. "github.com/onsi/gomega"
)

func TestMatchers(t *testing.T) {
	t.Parallel()

	matchTestCases := []struct {
		name        string
		input       any
		matcher     types.GomegaMatcher
		shouldFail  bool
		expectedMsg string
	}{
		// BePresent
		{
			name:       "BePresent succeeds for present lazy optional",
			input:      lazyoptional.Some("hello"),
			matcher:    matchers.BePresent[string](),
			shouldFail: false,
		},
		{
			name:        "BePresent fails for empty lazy optional",
			input:       lazyoptional.Empty[string](),
			matcher:     matchers.BePresent[string](),
			shouldFail:  true,
			expectedMsg: matchers.BePresentFailureMessage,
		},
		{
			name:       "Not(BePresent) succeeds for empty lazy optional",
			input:      lazyoptional.Empty[string](),
			matcher:    Not(matchers.BePresent[string]()),
			shouldFail: false,
		},
		{
			name:        "Not(BePresent) fails for present lazy optional",
			input:       lazyoptional.Some("hello"),
			matcher:     Not(matchers.BePresent[string]()),
			shouldFail:  true,
			expectedMsg: matchers.NotBePresentFailureMessage,
		},
		// BeEmpty
		{
			name:       "BeEmpty succeeds for empty lazy optional",
			input:      lazyoptional.Empty[string](),
			matcher:    matchers.BeEmpty[string](),
			shouldFail: false,
		},
		{
			name:        "BeEmpty fails for present lazy optional",
			input:       lazyoptional.Some("hello"),
			matcher:     matchers.BeEmpty[string](),
			shouldFail:  true,
			expectedMsg: matchers.BeEmptyFailureMessage,
		},
		{
			name:       "Not(BeEmpty) succeeds for present lazy optional",
			input:      lazyoptional.Some("hello"),
			matcher:    Not(matchers.BeEmpty[string]()),
			shouldFail: false,
		},
		{
			name:        "Not(BeEmpty) fails for empty lazy optional",
			input:       lazyoptional.Empty[string](),
			matcher:     Not(matchers.BeEmpty[string]()),
			shouldFail:  true,
			expectedMsg: matchers.NotBeEmptyFailureMessage,
		},
		// HaveValue
		{
			name:       "HaveValue succeeds for lazy optional with the same value",
			input:      lazyoptional.Some("hello"),
			matcher:    matchers.HaveValue("hello"),
			shouldFail: false,
		},
		{
			name:        "HaveValue fails for lazy optional with a different value",
			input:       lazyoptional.Some("world"),
			matcher:     matchers.HaveValue("hello"),
			shouldFail:  true,
			expectedMsg: matchers.HaveValueFailureMessage,
		},
		{
			name:        "HaveValue fails for an empty lazy optional",
			input:       lazyoptional.Empty[string](),
			matcher:     matchers.HaveValue("hello"),
			shouldFail:  true,
			expectedMsg: matchers.HaveValueFailureMessage,
		},
		{
			name:       "Not(HaveValue) succeeds for lazy optional with a different value",
			input:      lazyoptional.Some("world"),
			matcher:    Not(matchers.HaveValue("hello")),
			shouldFail: false,
		},
		{
			name:        "Not(HaveValue) fails for lazy optional with the same value",
			input:       lazyoptional.Some("hello"),
			matcher:     Not(matchers.HaveValue("hello")),
			shouldFail:  true,
			expectedMsg: matchers.NotHaveValueFailureMessage,
		},
	}

	for _, tc := range matchTestCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			if tc.shouldFail {
				failures := testutil.InterceptGomegaFailuresForTest(g, func() {
					g.Expect(tc.input).To(tc.matcher)
				})

				g.Expect(failures).To(HaveLen(1))
				g.Expect(failures[0]).To(ContainSubstring(tc.expectedMsg))
			} else {
				g.Expect(tc.input).To(tc.matcher)
			}
		})
	}

	errorTestCases := []struct {
		name        string
		matcher     types.GomegaMatcher
		input       any
		expectedErr error
	}{
		{
			name:        "BePresent returns error for non-lazy-optional",
			matcher:     matchers.BePresent[string](),
			input:       "not-a-lazy-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
		{
			name:        "BeEmpty returns error for non-lazy-optional",
			matcher:     matchers.BeEmpty[string](),
			input:       "not-a-lazy-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
		{
			name:        "HaveValue returns error for non-lazy-optional",
			matcher:     matchers.HaveValue("hello"),
			input:       "not-a-lazy-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
	}

	for _, tc := range errorTestCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := tc.matcher.Match(tc.input)

			g.Expect(err).To(MatchError(tc.expectedErr))
		})
	}
}

func TestFailureMessages(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	messages := []string{
		matchers.BePresentFailureMessage,
		matchers.NotBePresentFailureMessage,
		matchers.BeEmptyFailureMessage,
		matchers.NotBeEmptyFailureMessage,
		matchers.HaveValueFailureMessage,
		matchers.NotHaveValueFailureMessage,
	}

	for i, msg1 := range messages {
		for j, msg2 := range messages {
			if i == j {
				continue
			}
			g.Expect(strings.Contains(msg1, msg2)).To(BeFalse(), "failure messages should not be substrings of each other: '%s' contains '%s'", msg1, msg2)
		}
	}
}
