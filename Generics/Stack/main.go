package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}

	poppedElement := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return poppedElement, true
}

func main() {
	intStack := Stack[int]{}
	intStack.Push(12)
	intStack.Push(34)
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
}
