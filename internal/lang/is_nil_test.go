package lang_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/lang"

	. "github.com/onsi/gomega"
)

func TestReflectiveIsNil_WithUntypedNil(t *testing.T) {
	g := NewWithT(t)

	var input any

	actual := lang.ReflectiveIsNil(input)

	g.Expect(actual).To(BeTrue())
}

func TestReflectiveIsNil_WithTypedValues(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		name     string
		input    any
		expected bool
	}{
		{"nil pointer", (*int)(nil), true},
		{"nil interface", (any)(nil), true},
		{"nil slice", ([]int)(nil), true},
		{"nil map", (map[int]int)(nil), true},
		{"nil channel", (chan int)(nil), true},
		{"nil func", (func())(nil), true},
		{"non-nil int", 1, false},
		{"non-nil string", "hello", false},
		{"non-nil struct", struct{}{}, false},
		{"non-nil pointer", new(int), false},
		{"non-nil interface", any(1), false},
		{"non-nil slice", make([]int, 1), false},
		{"non-nil map", make(map[int]int), false},
		{"non-nil channel", make(chan int), false},
		{"non-nil func", func() {}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := lang.ReflectiveIsNil(tc.input)

			g.Expect(actual).To(Equal(tc.expected))
		})
	}
}
