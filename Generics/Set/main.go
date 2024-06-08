package main

import "fmt"

type Set[T comparable] struct {
	arr []T
}

func (s *Set[T]) AddElement(val T) {
	exists := false
	for _, value := range s.arr {
		if val == value {
			exists = true
			break
		}
	}
	if !exists {
		s.arr = append(s.arr, val)
	} else {
		fmt.Println(val, "is already in set")
	}
}

func main() {
	s := Set[int]{}
	s.AddElement(12)
	s.AddElement(23)
	s.AddElement(12)
	fmt.Println(s)
}
