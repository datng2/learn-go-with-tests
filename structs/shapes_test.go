package shapes

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   Rectangle{width: 1.0, height: 2.0},
			hasArea: 2.0,
		},
		{
			name:    "Circle",
			shape:   Circle{radius: 1.0},
			hasArea: 3.141592653589793,
		},
		{
			name:    "Triangle",
			shape:   Triangle{base: 1.0, height: 2.0},
			hasArea: 1.0,
		},
	}

	for _, test := range areaTests {
		// go test -run TestArea/Rectangle
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("For shape %#v , Got %f expected %f", test.shape, got, test.hasArea)
			}
		})

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
