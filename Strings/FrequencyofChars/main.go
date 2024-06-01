package main

import "fmt"

func main() {

	str := "GeeksForGeeks"
	m := make(map[rune]int)

	for _, val := range str {
		c := 0
		for _, v := range str {
			if val == v {
				c++
			}
		}
		m[val] = c
	}
	fmt.Println(m)
}
