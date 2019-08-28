package integer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter02(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integer Test Suite")
}
