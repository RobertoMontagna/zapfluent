package matchers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
	lazyoptionalmatchers "go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional/matchers"
)

var _ = Describe("LazyOptional Matchers", func() {
	Describe("BePresent", func() {
		When("the lazy optional is present", func() {
			It("succeeds", func() {
				Expect(lazyoptional.Some("hello")).Should(lazyoptionalmatchers.BePresent[string]())
			})
		})

		When("the lazy optional is empty", func() {
			It("fails", func() {
				Expect(lazyoptional.Empty[string]()).ShouldNot(lazyoptionalmatchers.BePresent[string]())
			})
		})

		When("the actual is not a lazy optional", func() {
			It("returns an error", func() {
				success, err := lazyoptionalmatchers.BePresent[string]().Match("not-a-lazy-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("BeEmpty", func() {
		When("the lazy optional is empty", func() {
			It("succeeds", func() {
				Expect(lazyoptional.Empty[string]()).Should(lazyoptionalmatchers.BeEmpty[string]())
			})
		})

		When("the lazy optional is present", func() {
			It("fails", func() {
				Expect(lazyoptional.Some("hello")).ShouldNot(lazyoptionalmatchers.BeEmpty[string]())
			})
		})

		When("the actual is not a lazy optional", func() {
			It("returns an error", func() {
				success, err := lazyoptionalmatchers.BeEmpty[string]().Match("not-a-lazy-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("HaveValue", func() {
		When("the lazy optional has the expected value", func() {
			It("succeeds", func() {
				Expect(lazyoptional.Some("hello")).Should(lazyoptionalmatchers.HaveValue("hello"))
			})
		})

		When("the lazy optional has a different value", func() {
			It("fails", func() {
				Expect(lazyoptional.Some("world")).ShouldNot(lazyoptionalmatchers.HaveValue("hello"))
			})
		})

		When("the lazy optional is empty", func() {
			It("fails", func() {
				Expect(lazyoptional.Empty[string]()).ShouldNot(lazyoptionalmatchers.HaveValue("hello"))
			})
		})

		When("the actual is not a lazy optional", func() {
			It("returns an error", func() {
				success, err := lazyoptionalmatchers.HaveValue("hello").Match("not-a-lazy-optional")
				Expect(success).To(BeFalse())
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
