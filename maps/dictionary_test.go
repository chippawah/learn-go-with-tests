package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("foo bar baz")
		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}
