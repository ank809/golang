package main

import "fmt"

func main() {
	age := 17
	if age >= 14 && age < 18 {
		fmt.Println("You can ride with parents!")
	} else if age >= 18 {
		fmt.Println("You can ride alone!")
	} else {
		fmt.Println("You cannot rise!")
	}
}
