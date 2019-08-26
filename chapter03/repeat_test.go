package repeat_test

import (
	"testing"

	"github.com/nsiregar/solid-go/chapter03"
)

func TestChapter03(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %q , expected: %q", result, expected)
		}
	}

	t.Run("repeat character 5 times", func(t *testing.T) {
		result := repeat.Repeat("a")
		expected := "aaaaa"

		assertEqual(t, result, expected)
	})
}

func BenchmarkChapter03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat.Repeat("a")
	}
}
