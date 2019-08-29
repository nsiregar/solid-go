package di_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter08/di"
)

var _ = Describe("Dependency Injection", func() {
	var buffer bytes.Buffer

	BeforeEach(func() {
		buffer = bytes.Buffer{}
	})

	Describe("#Greet", func() {
		It("write to buffer", func() {
			di.Greet(&buffer, "Chris")

			result := buffer.String()
			expected := "Hello, Chris"

			Expect(result).To(Equal(expected))
		})
	})
})
