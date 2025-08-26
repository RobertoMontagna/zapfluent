package testutil_test

import (
	"reflect"
	"testing"

	"go.robertomontagna.dev/zapfluent/testutil"

	. "github.com/onsi/gomega"
)

func TestInterceptGomegaFailures(t *testing.T) {
	g := NewWithT(t)

	tests := []struct {
		name              string
		f                 func(g *WithT)
		expectedToContain []any
	}{
		{
			name: "should return no failures when the function does not fail",
			f: func(g *WithT) {
				g.Expect(true).To(BeTrue())
			},
			expectedToContain: []any{},
		},
		{
			name: "should return a single failure when the function fails once",
			f: func(g *WithT) {
				g.Expect(true).To(BeFalse())
			},
			expectedToContain: []any{ContainSubstring("Expected\n    <bool>: true\nto be false")},
		},
		{
			name: "should return multiple failures when the function fails multiple times",
			f: func(g *WithT) {
				g.Expect(true).To(BeFalse())
				g.Expect(1).To(Equal(2))
			},
			expectedToContain: []any{
				ContainSubstring("Expected\n    <bool>: true\nto be false"),
				ContainSubstring("Expected\n    <int>: 1\nto equal\n    <int>: 2"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			failures := testutil.InterceptGomegaFailuresForTest(g, func() {
				tt.f(g)
			})

			g.Expect(failures).To(HaveLen(len(tt.expectedToContain)))
			for _, matcher := range tt.expectedToContain {
				g.Expect(failures).To(ContainElement(matcher))
			}
		})
	}
}

func TestInterceptGomegaFailures_RestoresFailHandler(t *testing.T) {
	g := NewWithT(t)
	originalFailHandler := g.Fail

	f := func() {
		g.Expect(true).To(BeFalse())
	}

	_ = testutil.InterceptGomegaFailuresForTest(g, f)

	originalFailHandlerPtr := reflect.ValueOf(originalFailHandler).Pointer()
	restoredFailHandlerPtr := reflect.ValueOf(g.Fail).Pointer()
	g.Expect(restoredFailHandlerPtr).To(Equal(originalFailHandlerPtr))
}
