package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {

		// when it fails the line number reported will be
		// in our function call rather than inside our test helper
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	engLang, plLang := "ENG", "PL"

	t.Run("saying hello to Mateusz in Polish", func(t *testing.T) {
		var name string = "Mateusz"
		got := Hello(name, plLang)
		want := "Witaj " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to Mateusz in English", func(t *testing.T) {
		var name string = "Mateusz"
		got := Hello(name, engLang)
		want := "Hello " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})

}
