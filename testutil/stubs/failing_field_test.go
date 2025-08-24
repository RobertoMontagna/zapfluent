package stubs_test

import (
	"errors"
	"testing"

	"go.robertomontagna.dev/zapfluent/testutil/stubs"

	. "github.com/onsi/gomega"
)

func TestFailingField_Name_ReturnsConfiguredName(t *testing.T) {
	g := NewWithT(t)

	field := stubs.FailingField{FieldName: "test_name"}

	name := field.Name()

	g.Expect(name).To(Equal("test_name"))
}

func TestFailingField_Name_ReturnsDefaultNameIfNotConfigured(t *testing.T) {
	g := NewWithT(t)

	field := stubs.FailingField{}

	name := field.Name()

	g.Expect(name).To(Equal("error"))
}

func TestFailingField_Encode_ReturnsConfiguredError(t *testing.T) {
	g := NewWithT(t)

	expectedErr := errors.New("encoding failed")
	field := stubs.FailingField{Err: expectedErr}

	err := field.Encode(nil)

	g.Expect(err).To(MatchError(expectedErr))
}

func TestNewFailingField(t *testing.T) {
	g := NewWithT(t)

	expectedName := "my-field"
	expectedErr := errors.New("my-error")

	field := stubs.NewFailingField(expectedName, expectedErr)

	g.Expect(field.FieldName).To(Equal(expectedName))
	g.Expect(field.Err).To(MatchError(expectedErr))
}
