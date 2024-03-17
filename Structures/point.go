package main

import "math"

type Point struct {
	x, y int
}

func distance(p1, p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y}
}

// METHOD : Associated with a specific datatype i.e struct
func (p1 Point) Distance(p2 Point) float64 {
	deltaX := p2.x - p1.x
	deltaY := p2.y - p1.y
	return math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY))
}
