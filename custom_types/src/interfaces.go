package src

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.Base + 2*t.Height
}

func ShapeArea(s Shape) float64 {
	return s.Area()
}

func ShapePerimeter(s Shape) float64 {
	return s.Perimeter()
}
