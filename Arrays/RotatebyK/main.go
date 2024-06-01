package main

import "fmt"

func main() {
	arr := []int{3, 5, 7, 9, 12, 6}
	k := 4

	for i := 0; i < k/2; i++ {
		temp := arr[i]
		arr[i] = arr[k-i-1]
		arr[k-i-1] = temp
	}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
