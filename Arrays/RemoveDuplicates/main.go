package main

import (
	"fmt"
	"sort"
)

func main() {
	ans := []int{}
	arr := []int{5, 9, 12, 3, 5, 6, 7, 6, 5, 6}
	// sort the array first

	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		if i == 0 || arr[i] != arr[i-1] {
			ans = append(ans, arr[i])
		}
	}
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}

}
