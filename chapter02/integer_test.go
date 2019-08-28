package integer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter02"
)

var _ = Describe("Integer", func() {
	Describe("Math operation", func() {
		Context("addition", func() {
			It("add two numbers", func() {
				result := integer.Add(1, 2)
				expected := 3

				Expect(result).To(Equal(expected))
			})
		})
	})
})
