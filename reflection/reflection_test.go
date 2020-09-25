package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("order guaranteed", func(t *testing.T) {

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
			{
				"Nested fields",
				Person{
					Name: "Chris",
					Profile: Profile{
						Age:  33,
						City: "London"}},
				[]string{"Chris", "London"},
			},
			{
				"Pointers to things",
				&Person{
					Name: "Matthew",
					Profile: Profile{
						Age:  33,
						City: "Poland"}},
				[]string{"Matthew", "Poland"},
			},
			{
				"Slices",
				[]Profile{
					Profile{Age: 33, City: "Norway"},
					Profile{Age: 34, City: "Spain"},
				},
				[]string{"Norway", "Spain"},
			},
			{
				"Arrays",
				[2]Profile{
					Profile{Age: 33, City: "Norway"},
					Profile{Age: 34, City: "Spain"},
				},
				[]string{"Norway", "Spain"},
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

	})

	t.Run("without order", func(t *testing.T) {

		testMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(testMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("channel test", func(t *testing.T) {
		channel := make(chan Profile)

		go func() {
			channel <- Profile{33, "Bar"}
			channel <- Profile{33, "Baz"}
			close(channel)
		}()

		var got []string
		want := []string{"Bar", "Baz"}

		walk(channel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", got, want)
		}

	})

	t.Run("func test", func(t *testing.T) {
		function := func() (Profile, Profile) {
			return Profile{33, "Bar"}, Profile{33, "Baz"}
		}

		var got []string
		want := []string{"Bar", "Baz"}

		walk(function, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", got, want)
		}

	})
}

func assertContains(t *testing.T, actual []string, expected string) {
	t.Helper()
	contains := false
	for _, element := range actual {
		if element == expected {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("Expected %s but not found", expected)
	}
}
