package main

type Shapes interface {
	area() float64
}

type Square struct {
	sidelength float64
}
type Triangle struct {
	base   float64
	height float64
}

func (s Square) area() float64 {
	return s.sidelength * s.sidelength
}

func (t Triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s Shapes) float64 {
	return s.area()
}
