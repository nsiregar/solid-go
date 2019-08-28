package repeat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter03"
)

var _ = Describe("Repeating Character", func() {
	Measure("running efficiently", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			output := repeat.Repeat("a")
			expected := "aaaaa"

			Expect(output).To(Equal(expected))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.2), "should not take long")
	}, 10000)
})
