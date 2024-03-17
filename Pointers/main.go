package main

import "fmt"

func main() {

	/*
		'*' is used for pointer dereferencing (accessing the value a pointer points to) and for declaring pointer types.
		'&' is used to get the memory address of a variable, returning a pointer to that variable.
	*/
	x := 10
	// Assigned y to address of x
	y := &x // it returns a pointer to here y is the pointer
	fmt.Println(x, y)
	// dereferenceing pointer
	*y = 7
	fmt.Println(x, y)
	str1, str2 := "Hello", "Hi"
	var pointer *string = &str1
	fmt.Println(str1, str2)
	changeValue(pointer, &str2)
	fmt.Println(str1, str2)

	var slice []int = []int{1, 2, 3, 4, 5}
	append_slice(&slice)
	fmt.Println(slice)

	fmt.Println(*(max_element(&slice)))

	fmt.Println(*(reverse_slices(&slice)))

}

func changeValue(str1, str2 *string) {
	temp := *str1
	*str1 = *str2
	*str2 = temp
}
