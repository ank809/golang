package main

import "fmt"

type Circle struct {
	radius int
}

func (c Circle) area() float64 {
	return (3.14 * float64(c.radius) * float64(c.radius))
}

type Rectangle struct {
	l, b int
}

func (r Rectangle) area() float64 {
	return float64(r.l * r.b)
}

type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	r := Rectangle{3, 5}
	c := Circle{3}
	fmt.Println(getArea(r))
	fmt.Println(getArea(c))

}
