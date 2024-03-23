package main

import "fmt"

func evenOrOdd() {
	fmt.Println("Function is odd or even called")
	var s []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range s {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
