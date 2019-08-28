package array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter04"
)

var _ = Describe("Array", func() {
	Describe("#Sum", func() {
		It("summarize array of 5 integers", func() {
			numbers := []int{1, 2, 3, 4, 5}

			result := array.Sum(numbers)
			expected := 15

			Expect(result).To(Equal(expected))
		})

		It("summarize array of n integers", func() {
			numbers := []int{1, 2, 3}

			result := array.Sum(numbers)
			expected := 6

			Expect(result).To(Equal(expected))
		})
	})

	Describe("#SumAll", func() {
		It("process array with #Sum", func() {
			result := array.SumAll([]int{1, 2}, []int{3, 4})
			expected := []int{3, 7}

			Expect(result).To(Equal(expected))
		})
	})

	Describe("#SumAllTails", func() {
		It("sum slice of array", func() {
			result := array.SumAllTails([]int{1, 2}, []int{3, 4})
			expected := []int{2, 4}

			Expect(result).To(Equal(expected))
		})

		It("works on empty array", func() {
			result := array.SumAllTails([]int{}, []int{3, 4})
			expected := []int{0, 4}

			Expect(result).To(Equal(expected))
		})
	})
})
