package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("It should return the string if the number is less than 1", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "a"

		assertCorrectString(t, repeated, expected)
	})
	t.Run("It should repeat the string the number of times given", func(t *testing.T) {
		repeated := Repeat("b", 2)
		expected := "bb"
		assertCorrectString(t, repeated, expected)
	})
}

func assertCorrectString(t *testing.T, repeated, expected string) {
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 5)
	fmt.Println(repeated)
	// Output: ccccc
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}
