package main

import (
	"fmt"
	"math"
)

func IsArmstrong() {
	var a int
	fmt.Println("Enter the number you want to check is Armstrong or not")
	fmt.Scan(&a)
	temp := a
	num := a
	length := 0

	for a > 0 {
		a /= 10
		length++
	}
	n := 0
	for temp > 0 {
		d := temp % 10
		n += int(math.Pow(float64(d), float64(length)))
		// fmt.Println(n)
		temp /= 10
	}

	if num == n {
		fmt.Printf("Number %d is armstrong", n)
	} else {
		fmt.Printf("Number %d is not armstrong", n)
	}
}
