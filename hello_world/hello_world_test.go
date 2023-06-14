package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to a specific person", func(t *testing.T) {
		got := Hello("YOU", "")
		want := "Hello, YOU!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Greet in Spanish", func(t *testing.T) {
		got := Hello("Carlos", "Spanish")
		want := "Hola, Carlos!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Greet in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjur, Pierre!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say hello to the world when given an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
