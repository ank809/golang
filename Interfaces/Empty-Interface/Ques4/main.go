package main

import "fmt"

func storeInSlice(items ...interface{}) {
	var s []interface{}
	s = append(s, items...)

	fmt.Println(s)
}

func main() {
	storeInSlice(12, "Ankita", 12.24)
	storeInSlice(12)
}
