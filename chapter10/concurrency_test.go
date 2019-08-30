package concurrency_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter10"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

var _ = Describe("Concurrency", func() {
	Describe("CheckWebsites", func() {
		It("returns boolean", func() {
			websites := []string{
				"http://google.com",
				"http://blog.gypsydave5.com",
				"waat://furhurterwe.geds",
			}

			expected := map[string]bool{
				"http://google.com":          true,
				"http://blog.gypsydave5.com": true,
				"waat://furhurterwe.geds":    false,
			}

			result := concurrency.CheckWebsites(mockWebsiteChecker, websites)

			Expect(result).To(Equal(expected))
		})
	})

	Measure("running concurrency", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			urls := make([]string, 100)
			for i := 0; i < len(urls); i++ {
				urls[i] = "a url"
			}

			concurrency.CheckWebsites(slowStubWebsiteChecker, urls)
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5), "should be fast")
	}, 1000)
})
