package shape_test

import (
	"testing"

	"github.com/nsiregar/solid-go/chapter05"
)

type Polygon interface {
	Area() float64
}

func TestChapter05Perimeter(t *testing.T) {
	assertEqual := func(t *testing.T, result, expected float64) {
		t.Helper()

		if result != expected {
			t.Errorf("result: %.2f , expected: %.2f", result, expected)
		}
	}

	t.Run("calculate rectangle perimeter", func(t *testing.T) {
		rectangle := shape.Rectangle{
			Width:  10.0,
			Height: 10.0,
		}

		result := rectangle.Perimeter()
		expected := 40.0

		assertEqual(t, result, expected)
	})
}

func TestChapter05Area(t *testing.T) {
	areaTests := []struct {
		name    string
		polygon Polygon
		area    float64
	}{
		{name: "Rectangle", polygon: shape.Rectangle{Width: 10.0, Height: 10.0}, area: 100.0},
		{name: "Circle", polygon: shape.Circle{Radius: 10.0}, area: 314.1592653589793},
		{name: "Triangle", polygon: shape.Triangle{Base: 12.0, Height: 6.0}, area: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.polygon.Area()

			if result != tt.area {
				t.Errorf("result: %.2f , expected: %.2f", result, tt.area)
			}
		})
	}
}
