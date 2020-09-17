package arraysandslices

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddSum(t *testing.T) {
	var given []int = []int{1, 2, 3, 4, 5}
	const expected = 15

	sum := AddSum(given)

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestArrSums(t *testing.T) {
	givenFirstArr := []int{1, 2, 3}
	givenSecArr := []int{4, 5}
	expected := []int{6, 9}

	sum := ArrSums(givenFirstArr, givenSecArr)

	if !cmp.Equal(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}

}
