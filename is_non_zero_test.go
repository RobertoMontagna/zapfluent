package zapfluent_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent"
)

func TestIsNotNil_WithUntypedNil(t *testing.T) {
	var input any = nil

	actual := zapfluent.IsNotNil(input)

	assert.False(t, actual)
}

func TestIsNotNil_WithTypedValues(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected bool
	}{
		{"nil pointer", (*int)(nil), false},
		{"nil interface", (interface{})(nil), false},
		{"nil slice", ([]int)(nil), false},
		{"nil map", (map[int]int)(nil), false},
		{"nil channel", (chan int)(nil), false},
		{"nil func", (func())(nil), false},
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
			actual := zapfluent.IsNotNil(tc.input)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
