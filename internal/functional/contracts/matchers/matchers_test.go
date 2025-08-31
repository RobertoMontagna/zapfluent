package matchers_test

import (
	"testing"

	"github.com/onsi/gomega/types"

	"go.robertomontagna.dev/zapfluent/internal/functional/contracts/matchers"
	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	"go.robertomontagna.dev/zapfluent/internal/testutil"

	. "github.com/onsi/gomega"
)

func TestMatchers(t *testing.T) {
	// This test table is designed to test the generic matchers with both
	// optional.Optional and lazyoptional.LazyOptional to ensure the
	// OptionalLike[T] interface is correctly handled.
	matchTestCases := []struct {
		name        string
		input       any
		matcher     types.GomegaMatcher
		shouldFail  bool
		expectedMsg string
	}{
		// BePresent with optional.Optional
		{
			name:       "BePresent succeeds for present optional",
			input:      optional.Some("hello"),
			matcher:    matchers.BePresent[string](),
			shouldFail: false,
		},
		{
			name:        "BePresent fails for empty optional",
			input:       optional.Empty[string](),
			matcher:     matchers.BePresent[string](),
			shouldFail:  true,
			expectedMsg: matchers.BePresentFailureMessage,
		},
		// BePresent with lazyoptional.LazyOptional
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
		// BeEmpty with optional.Optional
		{
			name:       "BeEmpty succeeds for empty optional",
			input:      optional.Empty[string](),
			matcher:    matchers.BeEmpty[string](),
			shouldFail: false,
		},
		{
			name:        "BeEmpty fails for present optional",
			input:       optional.Some("hello"),
			matcher:     matchers.BeEmpty[string](),
			shouldFail:  true,
			expectedMsg: matchers.BeEmptyFailureMessage,
		},
		// BeEmpty with lazyoptional.LazyOptional
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
		// HaveValue with optional.Optional
		{
			name:       "HaveValue succeeds for optional with the same value",
			input:      optional.Some("hello"),
			matcher:    matchers.HaveValue("hello"),
			shouldFail: false,
		},
		{
			name:        "HaveValue fails for an empty optional",
			input:       optional.Empty[string](),
			matcher:     matchers.HaveValue("hello"),
			shouldFail:  true,
			expectedMsg: matchers.HaveValueFailureMessage,
		},
		// HaveValue with lazyoptional.LazyOptional
		{
			name:       "HaveValue succeeds for lazy optional with the same value",
			input:      lazyoptional.Some("hello"),
			matcher:    matchers.HaveValue("hello"),
			shouldFail: false,
		},
		{
			name:        "HaveValue fails for an empty lazy optional",
			input:       lazyoptional.Empty[string](),
			matcher:     matchers.HaveValue("hello"),
			shouldFail:  true,
			expectedMsg: matchers.HaveValueFailureMessage,
		},
	}

	for _, tc := range matchTestCases {
		t.Run(tc.name, func(t *testing.T) {
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

	// Error handling tests
	errorTestCases := []struct {
		name        string
		matcher     types.GomegaMatcher
		input       any
		expectedErr error
	}{
		{
			name:        "BePresent returns error for non-optional type",
			matcher:     matchers.BePresent[string](),
			input:       "not-an-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
		{
			name:        "BeEmpty returns error for non-optional type",
			matcher:     matchers.BeEmpty[string](),
			input:       "not-an-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
		{
			name:        "HaveValue returns error for non-optional type",
			matcher:     matchers.HaveValue("hello"),
			input:       "not-an-optional",
			expectedErr: matchers.ErrMatcherWrongType,
		},
	}

	for _, tc := range errorTestCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			_, err := tc.matcher.Match(tc.input)

			g.Expect(err).To(MatchError(tc.expectedErr))
		})
	}
}
