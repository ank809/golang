package main

import "fmt"

func advancefunctions() {

	test := func(x int) {
		fmt.Println(-x)
	}
	test(7)

	// Calling the function at the time of declaratio
	test1 := func(x int) int {
		return -x
	}(7)
	fmt.Println(test1)

	ans := func(y int) int {
		return y * 2
	}
	fmt.Println(ans(3))
}
