package Shaper

type Shape interface {
	Area() float32
	Perimeter() float32
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return c.Radius * c.Radius * 3.14
}

func (c Circle) XXX() float32 {
	return c.Radius * c.Radius * 2.0
}

type Rectangle struct {
	Lendth, Width float32
}

func (r Rectangle) Area() float32 {
	return r.Lendth * r.Width
}
