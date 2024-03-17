package main

import "fmt"

func main() {

	// Int is an immutable datatype and its content cannot be changed and if you want to do so then you
	// need to assign a new value to variable holding the int.
	var x int = 12
	y := x
	y = 4
	fmt.Println(x, y)

	// Mutable datatypes : Maps, Slices
	// 1. Slice
	var s1 []int = []int{1, 2, 3, 5, 6}
	s2 := s1
	s2[0] = 100
	fmt.Println(s1, s2)
	// Here if we make any changes in s2 then those will be reflected back in s1 as s2 is just a reference to s1

	// 2. Maps

	var mp1 map[string]int = map[string]int{"helllo": 10, "seven": 7}
	mp2 := mp1
	mp2["eight"] = 5
	mp1["nine"] = 9
	fmt.Println(mp1, mp2)

	// Arrays

	var arr [4]int = [4]int{1, 2, 3, 4}
	arr1 := arr
	arr1[0] = 100
	fmt.Println(arr, arr1)
	// We cannot resize the array but we can actually modify individual elements.

	var slice []int = arr[0:2]
	slice[0] = 11
	fmt.Println(slice, arr)

}
