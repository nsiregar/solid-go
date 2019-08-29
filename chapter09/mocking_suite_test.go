package mocking_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestChapter09(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mocking Test Suite")
}
