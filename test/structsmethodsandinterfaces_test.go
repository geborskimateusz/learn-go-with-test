package test

import (
	"testing"

	geometry "geborskimateusz.com/m/structsmethodsandinterfaces"
)

func TestPerimeter(t *testing.T) {
	got := geometry.Perimeter(10.00, 20.00)
	want := 60.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := geometry.Area(10.00, 20.00)
	want := 200.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
