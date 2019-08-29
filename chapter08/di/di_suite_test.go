package di_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter08(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dependency Injection Test Suite")
}
