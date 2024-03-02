package main

import "fmt"

func advancefunctions() {

	test := func(x int) {
		fmt.Println(-x)
	}
	test(7)

	ans := func(y int) int {
		return y * 2
	}
	fmt.Println(ans(3))
}
