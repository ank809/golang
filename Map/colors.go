package main

import "fmt"

func colors() {
	color := map[string]int{
		"red":    1,
		"yellow": 2,
		"green":  3,
		"blue":   4,
	}
	fmt.Println(color)
	color["red"] = 9
	fmt.Println(color)
	delete(color, "red")
	// fmt.Println(color)
	printColors(color)

}

func printColors(c map[string]int) {
	c["blue"] = 12
	for color, value := range c {
		fmt.Println("Value for ", color, " is ", value)
	}
}
