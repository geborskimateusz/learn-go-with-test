package test

import (
	"testing"

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
	t.Skip("\nTO DO" +
		"\nCase 1: SumAll([]int{1,2}, []int{0,9}) would return []int{3, 9}i" +
		"\nCase 2: SumAll([]int{1,1,1}) would return []int{3}")
}
