package main

import "fmt"

func main() {

	fmt.Println("Maps")
	/* Maps are unordered collections of key-value pairs.
	Maps store data in key-value pairs and they do not care about the order.
	Maps are fast as comapred to array because in array we store element in order.
	Maps are reference type so we can modify them directly when we pass then into another functions.
	*/

	mp := map[string]int{
		"apple":  1,
		"mango":  5,
		"orange": 9,
	}
	fmt.Println(mp)

	var mp1 map[string]int = map[string]int{
		"apple":  1,
		"mango":  5,
		"orange": 9,
	}
	fmt.Println(mp1)

	// Access a particular value at key

	fmt.Println(mp["apple"])

	// To delete a value

	delete(mp, "apple")
	fmt.Println(mp)

	// To check if a key exist in map or not

	// Here value will store the value at that key and ok will store bool value that the key exists or not
	value, ok := mp1["tim"]
	fmt.Println(value, ok)

	colors()
}
