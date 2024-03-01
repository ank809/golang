package main

import "fmt"

func main() {

	x := 0
	// ----- ALWAYS TRUE ------

	//for {
	// 	fmt.Printf("Always True")
	// }

	for x < 5 {
		fmt.Println(x)
		x += 1
	}

	for x = 1; x < 5; x++ {
		fmt.Println(x)
	}
}
