package array_test

import (
	"reflect"
	"testing"

	"github.com/nsiregar/solid-go/chapter04"
)

func TestChapter04(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected int) {
		t.Helper()
		if result != expected {
			t.Errorf("result: %q , expected: %q", result, expected)
		}
	}

	t.Run("summarize array of 5 integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := array.Sum(numbers)
		expected := 15

		assertEqual(t, result, expected)
	})

	t.Run("summarize array of any size integers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		result := array.Sum(numbers)
		expected := 6

		assertEqual(t, result, expected)
	})
}

func TestChapter04PartTwo(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result: %v , expected: %v", result, expected)
		}
	}

	t.Run("sumAll array", func(t *testing.T) {
		result := array.SumAll([]int{1, 2}, []int{3, 4})
		expected := []int{3, 7}

		assertEqual(t, result, expected)
	})
}

func TestChapter04PartThree(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result: %v , expected: %v", result, expected)
		}
	}

	t.Run("sumAllTails sum some slice of array", func(t *testing.T) {
		result := array.SumAllTails([]int{1, 2}, []int{3, 4})
		expected := []int{2, 4}

		assertEqual(t, result, expected)
	})

	t.Run("sumAllTails sum some slice of empty array", func(t *testing.T) {
		result := array.SumAllTails([]int{}, []int{3, 4})
		expected := []int{0, 4}

		assertEqual(t, result, expected)
	})
}
