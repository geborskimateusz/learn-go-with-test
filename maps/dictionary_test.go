package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("on known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		given := "test"
		got, _ := dictionary.Search(given)
		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("throws ErrKeyNotFound", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		given := "unknown"
		_, err := dictionary.Search(given)
		want := ErrKeyNotFound

		asssertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add to a map", func(t *testing.T) {
		givenKey := "test"
		givenVal := "this is just a test"

		dictionary := Dictionary{}

		dictionary.Add(givenKey, givenVal)

		got, _ := dictionary.Search(givenKey)
		assertStrings(t, got, givenVal)
	})

	t.Run("throws ErrDuplicateKey", func(t *testing.T) {
		givenKey := "test"
		givenVal := "this is just a test"

		dictionary := Dictionary{}

		dictionary.Add(givenKey, givenVal)

		err := dictionary.Add(givenKey, givenVal)

		asssertError(t, err, ErrDuplicateKey)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func asssertError(t *testing.T, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected error got nil")
	}

	assertStrings(t, ErrKeyNotFound.Error(), want.Error())
}
