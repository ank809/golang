package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IsPalindrome() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter number you want to check is palindrome or not: ")
	scanner.Scan()
	num, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	a := num
	rev := 0
	for num > 0 {
		d := num % 10
		rev = rev*10 + int(d)
		num /= 10
	}
	if rev == int(a) {
		fmt.Printf("Number %d is palindrome \n", a)
	} else {
		fmt.Printf("Number %d is not palindrome\n", a)
	}
}
