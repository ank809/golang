package main

import (
	"fmt"
	"strconv"
)

func converToString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return strconv.Itoa(val)

	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", val)
	}
}

func main() {
	fmt.Println(converToString("ankita"))
	fmt.Printf("Value is %v and type is %T \n ", converToString(12), converToString(12))
	fmt.Println(converToString(13.23))
}
