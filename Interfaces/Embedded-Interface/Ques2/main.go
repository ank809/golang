package main

import "fmt"

type Drawable interface {
	Draw()
}
type Clickable interface {
	Click()
}

type Both interface {
	Drawable
	Clickable
}
type Button struct {
}

func (b Button) Draw() {
	fmt.Println("Drawable")
}
func (b Button) Click() {
	fmt.Println("Clickable")
}

func HandleInteraction(b Both) {
	b.Click()
	b.Draw()
}

func main() {
	var b Button
	HandleInteraction(b)
}
