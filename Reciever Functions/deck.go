package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
	// for i := 0; i < len(d); i++ {
	// 	fmt.Println(i, d[i])
	// }
}
