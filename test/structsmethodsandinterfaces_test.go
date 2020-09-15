package test

import (
	"testing"

	geometry "geborskimateusz.com/m/structsmethodsandinterfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := geometry.Rectangle{Width: 10, Height: 20}
	got := geometry.Perimeter(rectangle)
	want := 60.00
	assertEqual(t, got, want)
}

func TestArea(t *testing.T) {

	t.Run("Calculate area of circle", func(t *testing.T) {
		circle := geometry.Circle{Radius: 20}
		want := 1256.00
		calculateArea(t, circle, want)
	})

	t.Run("Calculate area of rectangle", func(t *testing.T) {
		rectangle := geometry.Rectangle{Width: 20.00, Height: 20.00}
		want := 400.00
		calculateArea(t, rectangle, want)
	})
}

func TestAreaTableDriven(t *testing.T) {
	arrTests := []TestTableStruct{
		{geometry.Circle{Radius: 20}, 1256.00},
		{geometry.Triangle{Base: 2, Height: 5}, 5},
		{geometry.Rectangle{Width: 20.00, Height: 20.00}, 400.00},
	}

	for _, tt := range arrTests {
		calculateArea(t, tt.shape, tt.want)
	}
}

func assertEqual(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

func calculateArea(t *testing.T, s geometry.Shape, want float64) {
	assertEqual(t, s.Area(), want)
}

type TestTableStruct struct {
	shape geometry.Shape
	want  float64
}
