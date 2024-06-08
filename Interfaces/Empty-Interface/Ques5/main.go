package main

import "fmt"

func createMap(keys []string, values []interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	for i := 0; i < len(keys); i++ {
		m[keys[i]] = values[i]
	}
	return m
}

func main() {
	keys := []string{
		"Hello", "Hi", "Hey", "Heylo",
	}

	values := []interface{}{1, 12.34, true, "hi"}
	fmt.Println(createMap(keys, values))

}
