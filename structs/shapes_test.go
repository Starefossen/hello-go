package structs

import (
	"testing"
)

func testReturn(t *testing.T, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	testReturn(t, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		area  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, area: 72.0},
		{shape: Circle{Radius: 10}, area: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, area: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.area {
			t.Errorf("%T got %.2f want %.2f", tt.shape, got, tt.area)
		}
	}
}
