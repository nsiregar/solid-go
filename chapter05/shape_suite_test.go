package shape_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter05(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shape Test Suite")
}
