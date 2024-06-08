package main

import "fmt"

type Worker interface {
	Work()
}

type Engineer struct{}

type Doctor struct{}

func (e Engineer) Work() {
	fmt.Println("Engineer Works")

}

func (d Doctor) Work() {
	fmt.Println("Doctor Works")
}

// func Working(w Worker) {
// 	w.Work()
// }

func main() {
	var e Engineer
	var d Doctor
	e.Work()
	d.Work()
	// Working(e)
	// Working(d)
}
