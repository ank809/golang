package main

import "fmt"

func binarySearch(arr []int, s int, e, target int) bool {
	for s <= e {
		m := (s + e) / 2
		if arr[m] == target {
			return true
		} else if arr[m] < target {
			s++
		} else {
			e--
		}
	}
	return false
}

func main() {

	var arr = []int{5, 6, 9, 10, 19, 21}
	// binary search

	fmt.Println(binarySearch(arr, 0, len(arr), 1))

}
