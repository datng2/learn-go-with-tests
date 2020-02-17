package shapes

import "math"

// Shape is an interface for Rect and Circle
type Shape interface {
	Area() float64
}

// Rectangle is a abstraction of a rect
type Rectangle struct {
	width  float64
	height float64
}

// Circle is a abstraction of a cir
type Circle struct {
	radius float64
}

// Triangle is a abstraction of a triangle
type Triangle struct {
	base   float64
	height float64
}

// Area returns area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Area returns area of a circle
func (c Circle) Area() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

// Area returns area of a Triangle
func (t Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}
