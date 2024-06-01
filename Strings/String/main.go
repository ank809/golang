package main

import "fmt"

func main() {
	// var a string
	a := "Hello"
	b := "World"
	c := a + " " + b
	fmt.Println(a, " ", b, " ", c)

	// substring

	d := a[:4]
	fmt.Println(d)

	fmt.Println(a[0])
	for _, val := range a {
		fmt.Println(val)
	}
}
