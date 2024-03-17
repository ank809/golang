// Find maximum element in a slice using pointers

package main

import "fmt"

func max_element(slice *[]int) *int {
	fmt.Println("Maximum Element in slice is")
	max := (*slice)[0]
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] > max {
			max = ((*slice)[i])
		}
	}
	return &max
}
