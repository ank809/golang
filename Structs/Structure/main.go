package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	age  int
	city string
}

func main() {

	person := Person{
		name: "Srijan",
		age:  19,
		city: "Dharamshala",
	}
	fmt.Println(person)
	fmt.Println(unsafe.Sizeof(person))
	fmt.Println(unsafe.Sizeof(person.age))
}
