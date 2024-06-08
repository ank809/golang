package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Car struct {
	Name   string
	Engine string
}

type Bike struct {
	Name   string
	Engine string
}

func (c Car) Drive() {
	fmt.Println(c.Name, c.Engine)
}

func (b Bike) Drive() {
	fmt.Println(b.Name, b.Engine)
}

func StartJourney(v Vehicle) {
	v.Drive()
}

func main() {
	c := Car{Name: "Range Rover", Engine: "Petrol"}
	StartJourney(c)

	b := Bike{Name: "Royal Enfield", Engine: "Petrol"}
	StartJourney(b)
}
