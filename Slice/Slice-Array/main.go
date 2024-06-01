package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4}
	slice := arr[:3]
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(slice, len(slice), cap(slice))

}
