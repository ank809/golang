package main

import "fmt"

func main() {
	fmt.Println("Slices in Go")

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	/*
		Slices are a part of array
		Slices are references to a contiguous segment of an array.
		Slices in Go works on top of array
		In array the length and capacity are same but in Slice it's not
	*/
	y := []int{1, 2, 3, 4, 5, 6}
	var s []int = a[1:3]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// We can append elements in a slice. It  do not modify the original one it creates a new one.
	s = append(s, 10)
	fmt.Println(s)

	fmt.Println(y)
	evenOrOdd()

}
