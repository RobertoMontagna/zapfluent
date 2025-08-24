package stubs_test

import (
	"errors"
	"testing"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/testutil/stubs"

	. "github.com/onsi/gomega"
)

var (
	errTest = errors.New("test error")
)

func TestFailingField_NewFailingField(t *testing.T) {
	g := NewWithT(t)
	fieldName := "test-field"

	result := stubs.NewFailingField(fieldName, errTest)

	g.Expect(result.FieldName).To(Equal(fieldName))
	g.Expect(result.Err).To(MatchError(errTest))
}

func TestFailingField_Encode(t *testing.T) {
	g := NewWithT(t)
	sut := stubs.NewFailingField("test-field", errTest)

	result := sut.Encode(zapcore.NewMapObjectEncoder())

	g.Expect(result).To(MatchError(errTest))
}

func TestFailingField_Name(t *testing.T) {
	g := NewWithT(t)

	tests := []struct {
		name         string
		sut          stubs.FailingField
		expectedName string
	}{
		{
			name:         "should return the configured name",
			sut:          stubs.NewFailingField("test-field", nil),
			expectedName: "test-field",
		},
		{
			name:         "should return 'error' if the name is empty",
			sut:          stubs.NewFailingField("", nil),
			expectedName: "error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.sut.Name()

			g.Expect(result).To(Equal(tt.expectedName))
		})
	}
}
