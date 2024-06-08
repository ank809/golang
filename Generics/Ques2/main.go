package main

import "fmt"

type Pair[T any] struct {
	a T
	b T
}

func (p *Pair[T]) Swap() {
	// var temp T
	temp := p.a
	p.a = p.b
	p.b = temp
}

func main() {

	p := Pair[int]{a: 12, b: 18}
	fmt.Println(p)
	p.Swap()
	fmt.Println(p)

}
