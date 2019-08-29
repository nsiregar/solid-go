package mocking_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
	"github.com/nsiregar/solid-go/chapter09"
)

var _ = Describe("Mocking", func() {
	Describe("#Countdown", func() {
		It("return current countdown", func() {
			buffer := &bytes.Buffer{}
			spySleeper := &mocking.SpySleeper{}

			mocking.Countdown(buffer, spySleeper)

			result := buffer.String()
			expected := `3
2
1
Go!`

			Expect(result).To(Equal(expected))
			Expect(spySleeper.Calls).To(Equal(4))
		})
	})
})
