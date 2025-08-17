package lazy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/lazy"
)

func TestNewConstantProducer(t *testing.T) {
	producer := lazy.NewConstantProducer("hello", 42)
	v1, v2 := producer()
	assert.Equal(t, "hello", v1)
	assert.Equal(t, 42, v2)
}
