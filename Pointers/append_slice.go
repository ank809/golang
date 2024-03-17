package main

import "fmt"

func append_slice(slice *[]int) {
	fmt.Println("Appending element in slice using pointers")
	fmt.Println(*slice)
	*slice = append(*slice, 100)
}
