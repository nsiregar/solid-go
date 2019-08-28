package hello_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter01(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hello Test Suite")
}
