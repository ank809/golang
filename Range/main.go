package main

import (
	"fmt"
)

func main() {
	fmt.Println("Use of Range in Array")

	a := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// Here i represent the index and element represent the element at that index
	for i, element := range a {
		fmt.Println(i, "corresponding element is ", element)
	}

	//If you do not  want to print index then use _ as it represnt anonymous var and you can't access it
	for _, element := range a {
		fmt.Println(element)
	}

	Duplicates()
}
