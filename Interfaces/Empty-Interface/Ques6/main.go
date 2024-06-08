package main

import "fmt"

func Sum(items ...interface{}) interface{} {
	sum := 0
	for _, item := range items {
		switch val := item.(type) {
		case int:
			sum += val
		case string:
			return val
		default:
			return "Nothing"
		}
	}
	return sum
}

func main() {
	fmt.Println(Sum(12, 34, 10))
	fmt.Println(Sum(12.23))
}
