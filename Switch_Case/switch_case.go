package main

import "fmt"

func main() {
	ans := -1

	switch ans {
	case 1, -1:
		fmt.Println("one")

	case 2, -2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("default")
	}
}
