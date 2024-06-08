package main

import "fmt"

func Describe(v interface{}) {
	fmt.Printf("Type is %T and value is %v \n", v, v)
}

type a struct {
	val int
}

func main() {
	s := a{val: 12}
	Describe(12)
	Describe("Hello")
	Describe(12.34)
	Describe(s)
}
