package shapes

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape        Shape
		expectedArea float64
	}{
		{shape: Rectangle{width: 1.0, height: 2.0}, expectedArea: 2.0},
		{shape: Circle{radius: 1.0}, expectedArea: 3.141592653589793},
		{shape: Triangle{base: 1.0, height: 2.0}, expectedArea: 2.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.expectedArea {
			t.Errorf("For shape %#v , Got %f expected %f", test.shape, got, test.expectedArea)
		}
	}
	// checkAreaFn := func(t *testing.T, shape Shape, expectedArea float64) {
	// 	area := shape.Area()
	// 	if area != expectedArea {
	// 		t.Errorf("Got %f expected %f", area, expectedArea)
	// 	}
	// }

	// t.Run("Area of a rectangle", func(t *testing.T) {
	// 	rect := Rectangle{1.0, 2.0}
	// 	checkAreaFn(t, rect, 2.0)
	// })

	// t.Run("Area of a circle", func(t *testing.T) {
	// 	circle := Circle{1.0}
	// 	checkAreaFn(t, circle, 3.141592653589793)
	// })
}
