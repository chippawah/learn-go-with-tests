package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Foo"
	var got []string
	x := struct {
		Name string
	}{expected}
	walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("wrong number of fn calls, got %d want 1", len(got))
	}
	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
