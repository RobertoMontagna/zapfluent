package lazy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.robertomontagna.dev/zapfluent/lazy"
)

func TestNewConstantProducer(t *testing.T) {
	expectedV1 := "hello"
	expectedV2 := 42

	producer := lazy.NewConstantProducer(expectedV1, expectedV2)
	v1, v2 := producer()

	assert.Equal(t, expectedV1, v1)
	assert.Equal(t, expectedV2, v2)
}
