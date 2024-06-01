package main

import "fmt"

/*
	Slices are a part of array
	Slices are references to a contiguous segment of an array.
	Slices in Go works on top of array
	In array the length and capacity are same but in Slice it's not
	If we assign one slice to another then the slices shares the undelying elements but if the cap and len of one slice changes it will not affect the other one
*/

func main() {
	s1 := []int{2, 3, 4, 5, 6}
	s2 := s1
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	s2 = append(s2, 12)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	s2[0] = 12
	s1[1] = 9
	fmt.Println(s1)
	fmt.Println(s2)
}
