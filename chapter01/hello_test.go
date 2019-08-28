package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter01"
)

var _ = Describe("Hello World", func() {
	Describe("Greeting people", func() {
		Context("with empty String", func() {
			It("should greet world", func() {
				result := hello.Greet("", "")
				expected := "Hello, world"

				Expect(result).To(Equal(expected))
			})
		})

		Context("with default language", func() {
			It("should be English", func() {
				result := hello.Greet("Chris", "")
				expected := "Hello, Chris"

				Expect(result).To(Equal(expected))
			})
		})

		Context("with custom language", func() {
			It("return in Spanish", func() {
				result := hello.Greet("Elodie", "Spanish")
				expected := "Hola, Elodie"

				Expect(result).To(Equal(expected))
			})

			It("return in French", func() {
				result := hello.Greet("Elodie", "French")
				expected := "Bonjour, Elodie"

				Expect(result).To(Equal(expected))
			})
		})
	})
})
