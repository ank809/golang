package main

import "fmt"

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() T {
	if len(q.data) == 0 {
		var zero T
		return zero
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}
func main() {
	q := Queue[int]{}
	q.Push(12)
	q.Push(10)
	q.Push(90)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
