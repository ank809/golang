package main

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() string {
	return "Woof-Woof"
}
func (c Cat) Speak() string {
	return "Meow-Meow"
}

func MakeSound(a Animal) string {
	return a.Speak()
}
