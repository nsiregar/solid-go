package integer_test

import (
	"testing"

	"github.com/nsiregar/solid-go/chapter02"
)

func TestChapter02(t *testing.T) {
	assertMessage := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %q , expected: %q", result, expected)
		}
	}

	t.Run("add two number", func(t *testing.T) {
		result := integer.Add(1, 2)
		expected := 3

		assertMessage(t, result, expected)
	})
}
