package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Map[T any](s []T, a func(s []T) T) T {
	return a(s)
}

func Add[T constraints.Ordered](s []T) T {
	var ans T
	for _, val := range s {
		ans += val
	}
	return ans
}
func Minimum[T constraints.Ordered](s []T) T {
	var min T
	min = s[0]
	for _, val := range s {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	s := []int{1, 23, 45, 10, 98}
	s1 := []string{"HI", "H"}
	fmt.Println(Map(s, Add))
	fmt.Println(Map(s1, Add))
	fmt.Println(Map(s, Minimum))
	fmt.Println(Map(s1, Minimum))
}
