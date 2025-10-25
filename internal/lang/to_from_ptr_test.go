package lang_test

import (
	"testing"

	"go.robertomontagna.dev/zapfluent/internal/lang"

	. "github.com/onsi/gomega"
)

func TestToPtr_And_MustFromPtr(t *testing.T) {
	g := NewWithT(t)

	testCases := []struct {
		name        string
		sourceValue any
	}{
		{
			"int",
			123,
		},
		{
			"string",
			"hello",
		},
		{
			"bool",
			true,
		},
		{
			"struct",
			struct {
				ID   int
				Name string
			}{
				1, "Test",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ptrValue := lang.ToPtr(tc.sourceValue)
			referencedValue := lang.MustFromPtr(ptrValue)

			g.Expect(ptrValue).ToNot(BeNil())
			g.Expect(*ptrValue).To(Equal(tc.sourceValue))
			g.Expect(ptrValue).To(BeAssignableToTypeOf(&tc.sourceValue))
			g.Expect(referencedValue).To(Equal(tc.sourceValue))
		})
	}
}
