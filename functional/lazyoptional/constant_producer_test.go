package lazyoptional_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

func TestNewConstantProducer(t *testing.T) {
	// Arrange
	expectedV1 := "hello"
	expectedV2 := 42

	// Act
	producer := lazyoptional.NewConstantProducer(expectedV1, expectedV2)
	v1, v2 := producer()

	// Assert
	assert.Equal(t, expectedV1, v1)
	assert.Equal(t, expectedV2, v2)
}
