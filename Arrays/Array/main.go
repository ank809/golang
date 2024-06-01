package main

// If we assign one array to another then all the elements are copies to another array and the two arrays don't share elements. IT means change on one array will not
// reflect on other
import "fmt"

func main() {
	// var n int
	// fmt.Println("Enter the size of array ")
	// fmt.Scan(&n)
	// arr := [10]int{}

	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&arr[i])
	// }

	// fmt.Println("Array Elements are : ")
	// for i := 0; i < n; i++ {
	// 	fmt.Println(arr[i])
	// }

	var a string
	var b int
	fmt.Println("Enter value of a  and b")
	fmt.Scan(&a, &b)
	fmt.Println(a, " ", b)

}
