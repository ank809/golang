package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FibonacciSeries() {
	fmt.Println("Fibonacci Series")
	fmt.Println("Enter the number")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	a := 0
	b := 1
	i := 1
	fmt.Println(a)
	fmt.Println(b)
	for i < int(num) {
		c := a + b
		fmt.Println(c)
		a = b
		b = c
		i++
	}
}
