package main

import "fmt"

func TwoDArray() {
	fmt.Println("Two Dimensional Array")

	// Static Initialization
	var a [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// Type Inference
	b := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(a[1][2])
	fmt.Println(b[0][1])
}
