package main

import "fmt"

func reverse_slices(slice *[]int) *[]int {
	fmt.Println("Reverse Slice using Pointers ")
	j := 0
	s1 := make([]int, len(*slice))
	for i := len(*slice) - 1; i >= 0; i-- {
		s1[j] = (*slice)[i]
		j++
	}
	return &s1
}
