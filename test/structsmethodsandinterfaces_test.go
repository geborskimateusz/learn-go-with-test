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
	circle := geometry.Circle{20}
	got := geometry.Area(circle)
	want := 1256.00

	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
