package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){

	// Ways to take input from the user
	// 1. Using Scan : In this new line  is considered as space.

	var a int
	fmt.Printf("Enter first number: ")
	fmt.Scan(&a)
	fmt.Println("Number you entered is ", a)

	// 2. Using Scanln 
	var num int
	fmt.Printf("Enter second number: ")
	fmt.Scanln(&num)
	fmt.Println("Number you entered is ", num)

	// 3. USing bufio

	//  Initializes a new Scanner object from the bufio package to read input from the standard input stream (os.Stdin)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter third number: ")
	scanner.Scan()
	c , _ := strconv.ParseInt(scanner.Text(), 10, 10)
	fmt.Println("Number is: ", c)

}