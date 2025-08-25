package matchers_test

import (
	"errors"
	"strings"
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional/matchers"
	"go.robertomontagna.dev/zapfluent/testutil"

	. "github.com/onsi/gomega"
)

// BePresent Tests
func TestMatcher_BePresent_SucceedsForPresent(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BePresent[string]()
	input := lazyoptional.Some("hello")

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_BePresent_FailsForEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BePresent[string]()
	input := lazyoptional.Empty[string]()

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.BePresentFailureMessage))
}

func TestMatcher_NotBePresent_SucceedsForEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.BePresent[string]())
	input := lazyoptional.Empty[string]()

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_NotBePresent_FailsForPresent(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.BePresent[string]())
	input := lazyoptional.Some("hello")

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.NotBePresentFailureMessage))
}

func TestMatcher_BePresent_ErrorForWrongType(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BePresent[string]()
	input := "not-a-lazy-optional"

	_, err := matcher.Match(input)

	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Is(err, matchers.ErrMatcherWrongType)).To(BeTrue())
}

// BeEmpty Tests
func TestMatcher_BeEmpty_SucceedsForEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BeEmpty[string]()
	input := lazyoptional.Empty[string]()

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_BeEmpty_FailsForPresent(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BeEmpty[string]()
	input := lazyoptional.Some("hello")

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.BeEmptyFailureMessage))
}

func TestMatcher_NotBeEmpty_SucceedsForPresent(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.BeEmpty[string]())
	input := lazyoptional.Some("hello")

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_NotBeEmpty_FailsForEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.BeEmpty[string]())
	input := lazyoptional.Empty[string]()

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.NotBeEmptyFailureMessage))
}

func TestMatcher_BeEmpty_ErrorForWrongType(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.BeEmpty[string]()
	input := "not-a-lazy-optional"

	_, err := matcher.Match(input)

	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Is(err, matchers.ErrMatcherWrongType)).To(BeTrue())
}

// HaveValue Tests
func TestMatcher_HaveValue_SucceedsForSameValue(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.HaveValue("hello")
	input := lazyoptional.Some("hello")

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_HaveValue_FailsForDifferentValue(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.HaveValue("hello")
	input := lazyoptional.Some("world")

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.HaveValueFailureMessage))
}

func TestMatcher_HaveValue_FailsForEmpty(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.HaveValue("hello")
	input := lazyoptional.Empty[string]()

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.HaveValueFailureMessage))
}

func TestMatcher_NotHaveValue_SucceedsForDifferentValue(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.HaveValue("hello"))
	input := lazyoptional.Some("world")

	success, err := matcher.Match(input)

	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(success).To(BeTrue())
}

func TestMatcher_NotHaveValue_FailsForSameValue(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := Not(matchers.HaveValue("hello"))
	input := lazyoptional.Some("hello")

	failures := testutil.InterceptGomegaFailuresForTest(g, func() {
		g.Expect(input).To(matcher)
	})

	g.Expect(failures).To(HaveLen(1))
	g.Expect(failures[0]).To(ContainSubstring(matchers.NotHaveValueFailureMessage))
}

func TestMatcher_HaveValue_ErrorForWrongType(t *testing.T) {
	g := NewGomegaWithT(t)

	matcher := matchers.HaveValue("hello")
	input := "not-a-lazy-optional"

	_, err := matcher.Match(input)

	g.Expect(err).To(HaveOccurred())
	g.Expect(errors.Is(err, matchers.ErrMatcherWrongType)).To(BeTrue())
}

// Other Tests
func TestFailureMessages_AreNotSubstrings(t *testing.T) {
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
			const m1 = "failure messages should not be substrings of each other"
			const m2 = "'%s' contains '%s'"
			g.Expect(strings.Contains(msg1, msg2)).To(BeFalse(), m1+": "+m2, msg1, msg2)
		}
	}
}
