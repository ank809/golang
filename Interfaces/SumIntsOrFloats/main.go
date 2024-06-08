package main

import (
	"fmt"
)

func SumIntsOrFloats(arr []interface{}) interface{} {
	var sumInt int
	var sumFloat float64
	intCount := 0
	floatCount := 0

	for _, v := range arr {
		switch v := v.(type) {
		case int:
			sumInt += v
			intCount++
		case float64:
			sumFloat += v
			floatCount++
		default:
			return "Unsupported type"
		}
	}

	if intCount > 0 && floatCount == 0 {
		return sumInt
	} else if floatCount > 0 && intCount == 0 {
		return sumFloat
	} else {
		return "Mixed types"
	}
}

func main() {
	ints := []interface{}{1, 2, 3, 4}
	floats := []interface{}{1.1, 2.2, 3.3}
	mixed := []interface{}{1, 2.2, 3}

	fmt.Println(SumIntsOrFloats(ints))   // Output: 10
	fmt.Println(SumIntsOrFloats(floats)) // Output: 6.6
	fmt.Println(SumIntsOrFloats(mixed))  // Output: Mixed types
}
