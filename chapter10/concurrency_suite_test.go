package concurrency_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter10(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Concurrency Test Suite")
}
