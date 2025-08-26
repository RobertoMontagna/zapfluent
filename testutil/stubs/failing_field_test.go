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

func TestNewFailingField(t *testing.T) {
	const fieldName = "test-field"

	testCases := []struct {
		name          string
		options       []stubs.FailingFieldForTestOption
		expectedName  string
		expectedError string
	}{
		{
			name:          "with default values",
			options:       []stubs.FailingFieldForTestOption{},
			expectedName:  "error",
			expectedError: "unspecified error",
		},
		{
			name: "with specified values",
			options: []stubs.FailingFieldForTestOption{
				stubs.WithName(fieldName),
				stubs.WithError(errTest),
			},
			expectedName:  fieldName,
			expectedError: "test error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			sut := stubs.NewFailingFieldForTest(tc.options...)

			g.Expect(sut.Name()).To(Equal(tc.expectedName))
			g.Expect(sut.Encode(zapcore.NewMapObjectEncoder())).To(MatchError(tc.expectedError))
		})
	}
}

func TestFailingFieldOptions_Panics(t *testing.T) {
	testCases := []struct {
		name        string
		f           func()
		expectedMsg string
	}{
		{
			name: "WithName panics if name is empty",
			f: func() {
				stubs.NewFailingFieldForTest(stubs.WithName(""))
			},
			expectedMsg: "name cannot be empty",
		},
		{
			name: "WithError panics if error is nil",
			f: func() {
				stubs.NewFailingFieldForTest(stubs.WithError(nil))
			},
			expectedMsg: "error cannot be nil",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)

			g.Expect(tc.f).To(PanicWith(tc.expectedMsg))
		})
	}
}
