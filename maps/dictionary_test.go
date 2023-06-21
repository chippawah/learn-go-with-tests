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

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defintion := "this is a test"
		dictionary.Add(word, defintion)
		assertDefinition(t, dictionary, word, defintion)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "new defintion")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word string, defintion string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, defintion)
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
