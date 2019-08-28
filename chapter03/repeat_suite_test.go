package repeat_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter03(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repeat Test Suite")
}
