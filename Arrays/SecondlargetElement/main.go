package main

import "fmt"

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func main() {
	arr := []int{12, 3, 45, 6, 23, 90}

	// find the second largest element by sorting the array
	sort(arr)
	a := len(arr)
	fmt.Println(arr)
	fmt.Println("Second largest element is ", arr[a-2])

}
