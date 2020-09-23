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
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Matthew", "London"},
			[]string{"Matthew", "London"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				Age  int
			}{"Matthew", 33},
			[]string{"Matthew"},
		},
	}

	for _, test := range cases {

		var got []string

		walk(test.Input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, test.Expected) {
			t.Errorf("got %v want %v", got, test.Expected)
		}
	}

}
