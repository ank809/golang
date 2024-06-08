package main

import "fmt"

func PrintVal[T any](val T) {
	fmt.Println(val)
}

func main() {
	PrintVal(12)
	PrintVal("Hello")
}
