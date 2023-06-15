package structs

import "testing"

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func TestPerimeter(t *testing.T) {
	checkResult := func(t *testing.T, got, want float64) {
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		checkResult(t, got, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Circumference()
		want := 62.83185307179586
		checkResult(t, got, want)
	})
}
