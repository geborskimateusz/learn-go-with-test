package test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	arraysAndSlices "geborskimateusz.com/m/arraysandslices"
)

func TestAddSum(t *testing.T) {
	var given []int = []int{1, 2, 3, 4, 5}
	const expected = 15

	sum := arraysAndSlices.AddSum(given)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestArrSums(t *testing.T) {
	givenFirstArr := []int{1, 2, 3}
	givenSecArr := []int{4, 5}
	expected := []int{6, 9}

	sum := arraysAndSlices.ArrSums(givenFirstArr, givenSecArr)

	if !cmp.Equal(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}

}
