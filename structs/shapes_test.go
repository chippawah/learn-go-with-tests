package structs

import "testing"

func TestShapes(t *testing.T) {
	checkResult := func(t *testing.T, got, want float64) {
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("calculate the permiter of a rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		checkResult(t, got, want)
	})
	t.Run("calulate the area of a rectangle", func(t *testing.T) {
		got := Area(12.0, 6.0)
		want := 72.0
		checkResult(t, got, want)
	})
}
