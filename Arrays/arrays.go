package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)

	// Integer array
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		scanner.Scan()
		arr[i], _ = strconv.Atoi(scanner.Text())
		sum = sum + arr[i]
	}
	fmt.Println("Sum of elements of Arrays", sum)

	// String array
	var str [5]string

	for j := 0; j < 5; j++ {
		scanner.Scan()
		str[j] = scanner.Text()
	}
	fmt.Println("Elements in string arrays are:")
	for j := 0; j < 5; j++ {
		fmt.Println(str[j])
	}

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	TwoDArray()
}
