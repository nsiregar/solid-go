package hello_test

import (
	"testing"

	"github.com/nsiregar/solid-go/chapter01"
)

func TestChapter01(t *testing.T) {
	assertMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %q , expected: %q", result, expected)
		}
	}

	t.Run("greeting hello to people", func(t *testing.T) {
		result := hello.Greet("Chris", "")
		expected := "Hello, Chris"

		assertMessage(t, result, expected)
	})

	t.Run("greeting hello to world, when empty string supplied", func(t *testing.T) {
		result := hello.Greet("", "")
		expected := "Hello, world"

		assertMessage(t, result, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		result := hello.Greet("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assertMessage(t, result, expected)
	})

	t.Run("in French", func(t *testing.T) {
		result := hello.Greet("Elodie", "French")
		expected := "Bonjour, Elodie"

		assertMessage(t, result, expected)
	})
}
