package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	fmt.Print("Enter two  numbers: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	c := a + b
	fmt.Println("Sum of numbers is ", c)

	d := a / b
	fmt.Println("a/b is ", d)

	fmt.Println("a % b is ", a%b)

	fmt.Println(math.Sqrt(12))

	fmt.Println(a == 0 && b == 5)

	fmt.Println(a & b)
}
