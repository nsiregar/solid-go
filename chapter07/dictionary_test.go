package dictionary_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter07"
)

var _ = Describe("Dictionary", func() {
	var map_struct dictionary.Dictionary

	BeforeEach(func() {
		map_struct = dictionary.Dictionary{"test": "test value"}
	})

	Describe("#Search", func() {
		It("search key and return value", func() {
			result, _ := map_struct.Search("test")
			expected := "test value"

			Expect(result).To(Equal(expected))
		})

		Context("undefined key", func() {
			It("raise error", func() {
				_, err := map_struct.Search("undefined")

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("#Add", func() {
		It("append key value to object", func() {
			map_struct.Add("key", "value")

			result, err := map_struct.Search("key")
			expected := "value"

			Expect(result).To(Equal(expected))
			Expect(err).NotTo(HaveOccurred())
		})

		Context("key exist", func() {
			It("raise error", func() {
				map_struct.Add("key", "value")
				err := map_struct.Add("key", "something else")

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("#Update", func() {
		It("update value of key", func() {
			map_struct.Update("test", "new value")

			result, err := map_struct.Search("test")
			expected := "new value"

			Expect(result).To(Equal(expected))
			Expect(err).NotTo(HaveOccurred())
		})

		Context("key not exist", func() {
			It("raise error", func() {
				err := map_struct.Update("hello", "new value")

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("#Delete", func() {
		It("delete key and value", func() {
			map_struct.Add("key", "value")
			map_struct.Delete("test")

			_, err := map_struct.Search("test")

			Expect(err).To(HaveOccurred())
		})

		Context("key not exist", func() {
			It("raise error", func() {
				err := map_struct.Delete("hello")

				Expect(err).To(HaveOccurred())
			})
		})
	})
})
