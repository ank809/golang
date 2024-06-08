package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length, Breadth float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

// Extra method in Rectangle
func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.Length*r.Length + r.Breadth*r.Breadth)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius * c.Radius
}

func Calculate(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
	// Simple type assertion
	// Panic if shape is not rectangle
	// fmt.Println(s.(Rectangle).Diagonal())

	if rec, ok := s.(Rectangle); ok {
		fmt.Println("Inside if ", rec.Diagonal())
	} else {
		fmt.Println("No diagonal method ")
	}

	switch shape := s.(type) {
	case Rectangle:
		fmt.Println("Inside switch", shape.Diagonal())
	case Circle:
		fmt.Println("Shape is circle so no diagonal method")
	default:
		fmt.Println("Other Shape")
	}

}

func main() {
	r := Rectangle{Length: 10, Breadth: 10}
	c := Circle{Radius: 4}
	Calculate(r)
	Calculate(c)
}
