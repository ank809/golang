// Print all the duplicates elements in array
package main

import "fmt"

func Duplicates() {
	a := []int{1, 2, 3, 4, 5, 6, 6, 4, 2}

	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]
			if element == element2 {
				fmt.Println("Duplicate element is", element)
			}
		}
	}
}
