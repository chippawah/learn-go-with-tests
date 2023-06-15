package structs

import "testing"

type tableTest struct {
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {
	areaTests := []tableTest{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
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
		got := circle.Perimeter()
		want := 62.83185307179586
		checkResult(t, got, want)
	})
}
