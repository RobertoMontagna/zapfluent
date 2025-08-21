package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
	optionalmatchers "go.robertomontagna.dev/zapfluent/internal/functional/optional/matchers"
)

var _ = Describe("Optional Matchers", func() {
	Describe("BePresent", func() {
		When("the optional is present", func() {
			It("succeeds", func() {
				Expect(optional.Some("hello")).Should(optionalmatchers.BePresent[string]())
			})
		})

		When("the optional is empty", func() {
			It("fails", func() {
				Expect(optional.Empty[string]()).ShouldNot(optionalmatchers.BePresent[string]())
			})
		})

		When("the actual is not an optional", func() {
			It("returns an error", func() {
				success, err := optionalmatchers.BePresent[string]().Match("not-an-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("BeEmpty", func() {
		When("the optional is empty", func() {
			It("succeeds", func() {
				Expect(optional.Empty[string]()).Should(optionalmatchers.BeEmpty[string]())
			})
		})

		When("the optional is present", func() {
			It("fails", func() {
				Expect(optional.Some("hello")).ShouldNot(optionalmatchers.BeEmpty[string]())
			})
		})

		When("the actual is not an optional", func() {
			It("returns an error", func() {
				success, err := optionalmatchers.BeEmpty[string]().Match("not-an-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("HaveValue", func() {
		When("the optional has the expected value", func() {
			It("succeeds", func() {
				Expect(optional.Some("hello")).Should(optionalmatchers.HaveValue("hello"))
			})
		})

		When("the optional has a different value", func() {
			It("fails", func() {
				Expect(optional.Some("world")).ShouldNot(optionalmatchers.HaveValue("hello"))
			})
		})

		When("the optional is empty", func() {
			It("fails", func() {
				Expect(optional.Empty[string]()).ShouldNot(optionalmatchers.HaveValue("hello"))
			})
		})

		When("the actual is not an optional", func() {
			It("returns an error", func() {
				success, err := optionalmatchers.HaveValue("hello").Match("not-an-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
