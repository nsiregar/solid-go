package shape_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nsiregar/solid-go/chapter05"
)

var _ = Describe("Shape", func() {
	Describe("#Perimeter", func() {
		It("calculate rectangle perimeter", func() {
			rectangle := shape.Rectangle{
				Width:  10.0,
				Height: 10.0,
			}

			result := rectangle.Perimeter()
			expected := 40.0

			Expect(result).To(Equal(expected))
		})
	})

	Describe("#Area", func() {
		It("calculate shape area", func() {
			areaTests := []struct {
				name  string
				shape shape.Shape
				area  float64
			}{
				{name: "Rectangle", shape: shape.Rectangle{Width: 10.0, Height: 10.0}, area: 100.0},
				{name: "Circle", shape: shape.Circle{Radius: 10.0}, area: 314.1592653589793},
				{name: "Triangle", shape: shape.Triangle{Base: 12.0, Height: 6.0}, area: 36.0},
			}

			for _, tt := range areaTests {
				result := tt.shape.Area()

				Expect(result).To(Equal(tt.area))
			}
		})
	})
})
