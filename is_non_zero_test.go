package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent"
)

func TestIsNotNil(t *testing.T) {
	var p *int
	var i interface{}
	var s []int
	var m map[int]int
	var c chan int
	var f func()

	testCases := []struct {
		name     string
		input    any
		expected bool
	}{
		{"untyped nil", nil, false},
		{"nil pointer", p, false},
		{"nil interface", i, false},
		{"nil slice", s, false},
		{"nil map", m, false},
		{"nil channel", c, false},
		{"nil func", f, false},
		{"non-nil int", 1, true},
		{"non-nil string", "hello", true},
		{"non-nil struct", struct{}{}, true},
		{"non-nil pointer", new(int), true},
		{"non-nil interface", interface{}(1), true},
		{"non-nil slice", make([]int, 1), true},
		{"non-nil map", make(map[int]int), true},
		{"non-nil channel", make(chan int), true},
		{"non-nil func", func() {}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == nil {
				assert.Equal(t, tc.expected, zapfluent.IsNotNil[any](nil))
			} else {
				assert.Equal(t, tc.expected, zapfluent.IsNotNil(tc.input))
			}
		})
	}
}
