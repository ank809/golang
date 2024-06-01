package main

import "fmt"

func main() {
	m := map[int]string{}

	m[0] = "Hello"
	m[2] = "Hey"
	fmt.Println(m)
	m1 := m
	m1[0] = "Hi"
	fmt.Println(m)
	fmt.Println(m1)

	s, exist := m[8]
	fmt.Println(exist, s)
	delete(m, 0)
	m[0] = "Heylooooo"
	fmt.Println(m)
}
