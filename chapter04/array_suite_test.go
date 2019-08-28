package array_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter04(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Array Test Suite")
}
