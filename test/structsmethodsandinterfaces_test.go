package test

import (
	"testing"

	geometry "geborskimateusz.com/m/structsmethodsandinterfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := geometry.Rectangle{10, 20}
	got := geometry.Perimeter(rectangle)
	want := 60.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertEqual := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	calculateArea := func(t *testing.T, s geometry.Shape, want float64) {
		assertEqual(t, s.Area(), want)
	}

	t.Run("Calculate area of circle", func(t *testing.T) {
		circle := geometry.Circle{20}
		want := 1256.00
		calculateArea(t, circle, want)
	})

	t.Run("Calculate area of rectangle", func(t *testing.T) {
		rectangle := geometry.Rectangle{20.00, 20.00}
		want := 400.00
		calculateArea(t, rectangle, want)
	})
}
