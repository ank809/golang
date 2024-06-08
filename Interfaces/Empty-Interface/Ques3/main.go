package main

import "fmt"

func process(v interface{}) interface{} {
	switch val := v.(type) {
	case string:
		// fmt.Println(val)
		val = "Go"
		return val
	case int:
		val += 1
		return val
	case bool:
		val = !val
		return val
	default:
		return "Nothing"
	}
}

func main() {
	fmt.Println(process("Ankita"))
	fmt.Println(process(true))
	fmt.Println(process(12))
}
