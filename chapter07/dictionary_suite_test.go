package dictionary_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter07(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dictionary Test Suite")
}
