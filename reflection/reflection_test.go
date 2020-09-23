package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		TestDescription string
		Input           interface{}
		Expected        []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Matthew"},
			[]string{"Matthew"},
		},
	}

	for _, test := range cases {

		var got []string

		walk(test.Input, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("Wrong number of function calls, expected %d got %d ", 1, len(got))
		}

		if !reflect.DeepEqual(got, test.Expected) {
			t.Errorf("got %v want %v", got, test.Expected)
		}
	}

}
