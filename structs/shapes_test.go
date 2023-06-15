package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tableTest := range areaTests {
		got := tableTest.shape.Area()
		if got != tableTest.want {
			t.Errorf("got %g want %g", got, tableTest.want)
		}
	}
}

// Here it doesn't make as much sense to do a table based test because the method names are different

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
