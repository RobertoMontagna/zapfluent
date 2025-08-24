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

func TestFailingField_NewFailingField_ShouldReturnFieldWithDefaultValues(t *testing.T) {
	g := NewWithT(t)
	sut := stubs.NewFailingFieldForTest()

	g.Expect(sut.Name()).To(Equal("error"))
	g.Expect(sut.Encode(zapcore.NewMapObjectEncoder())).To(MatchError("unspecified error"))
}

func TestFailingField_NewFailingField_ShouldReturnFieldWithSpecifiedValues(t *testing.T) {
	g := NewWithT(t)
	fieldName := "test-field"

	sut := stubs.NewFailingFieldForTest(
		stubs.WithName(fieldName),
		stubs.WithError(errTest),
	)

	g.Expect(sut.Name()).To(Equal(fieldName))
	g.Expect(sut.Encode(zapcore.NewMapObjectEncoder())).To(MatchError(errTest))
}

func TestFailingField_WithName_ShouldPanicIfNameIsEmpty(t *testing.T) {
	g := NewWithT(t)

	g.Expect(func() {
		stubs.NewFailingFieldForTest(stubs.WithName(""))
	}).To(PanicWith("name cannot be empty"))
}

func TestFailingField_WithError_ShouldPanicIfErrorIsNil(t *testing.T) {
	g := NewWithT(t)

	g.Expect(func() {
		stubs.NewFailingFieldForTest(stubs.WithError(nil))
	}).To(PanicWith("error cannot be nil"))
}

func TestFailingField_Encode(t *testing.T) {
	g := NewWithT(t)
	sut := stubs.NewFailingFieldForTest(
		stubs.WithName("test-field"),
		stubs.WithError(errTest),
	)

	result := sut.Encode(zapcore.NewMapObjectEncoder())

	g.Expect(result).To(MatchError(errTest))
}

func TestFailingField_Name(t *testing.T) {
	g := NewWithT(t)

	sut := stubs.NewFailingFieldForTest(
		stubs.WithName("test-field"),
		stubs.WithError(errTest),
	)

	result := sut.Name()

	g.Expect(result).To(Equal("test-field"))
}
