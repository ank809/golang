package main

import "fmt"

type Income interface {
	Source() string
	Calculate() int
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (f FixedBilling) Source() string {
	return f.projectName
}

func (f FixedBilling) Calculate() int {
	return f.biddedAmount
}

func (t TimeAndMaterial) Source() string {
	return t.projectName
}

func (t TimeAndMaterial) Calculate() int {
	return t.hourlyRate * t.noOfHours
}

func PrintVal(i []Income) {
	for _, income := range i {
		fmt.Println(income.Calculate(), income.Source())
	}
}

func main() {
	p := FixedBilling{projectName: "Project1", biddedAmount: 120000}
	t := TimeAndMaterial{projectName: "Project2", hourlyRate: 1500, noOfHours: 6}
	income := []Income{p, t}
	PrintVal(income)
}
