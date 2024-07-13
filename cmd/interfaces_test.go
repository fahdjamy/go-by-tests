package main

import "testing"

func TestShapeInterfaces(t *testing.T) {

	verifyAreaResp := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		response := shape.Area()
		if response != expected {
			t.Errorf("got %g, want %g", response, expected)
		}
	}

	t.Run("shape interfaces", func(t *testing.T) {
		circleShape := Circle{
			Radius: 10,
		}
		rectangleShape := Rectangle{
			Width:  10,
			Height: 10,
		}
		verifyAreaResp(t, rectangleShape, 40.00)
		verifyAreaResp(t, circleShape, 314.1592653589793)
	})

	t.Run("table testing shape interfaces", func(t *testing.T) {
		shapes := []struct {
			name  string
			shape Shape
			want  float64
		}{
			{"Circle", Circle{Radius: 5}, 78.53981633974483},
			{"Rectangle", Rectangle{Width: 10, Height: 10}, 40.00},
		}

		for _, shape := range shapes {
			t.Run(shape.name, func(t *testing.T) {
				verifyAreaResp(t, shape.shape, shape.want)
			})
		}
	})
}
