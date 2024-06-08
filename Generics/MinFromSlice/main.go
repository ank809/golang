package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func findMinimum[T constraints.Ordered](s []T) (min T) {
	if len(s) == 0 {
		panic("Empty slice ")
	}
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min

}

func main() {
	intSlice := []int{3, 2, 1, 5, 74, 3}
	fmt.Println(findMinimum(intSlice))
}
