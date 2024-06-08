package main

import "fmt"

type Animal interface {
	Speak() string
}
type Dog struct {
}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof-Woof"
}

func (c Cat) Speak() string {
	return "Meow -Meow "
}
func MakeSound(a Animal) string {
	return a.Speak()
}

func main() {
	d := Dog{}
	fmt.Println(MakeSound(d))
}
