package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

// This function cannot change the value of firstName because it operates on a copy of the Person struct.
// Therefore, despite the method being called, the change is not reflected in the original p variable.

func (p Person) changeFirstName(newName string) {
	p.firstName = newName
}

func (p *Person) changeName(newName string) {
	(*p).firstName = newName
}
func main() {
	p := Person{"Ankita", "Thakur"}
	fmt.Printf("%+v \n", p)
	p.changeFirstName("Astha")
	fmt.Printf("%+v \n", p)
	person := &p
	p.changeName("Anshuman")
	fmt.Printf("%+v \n", p)
	person.changeName("Astha")
	fmt.Printf("%+v \n", p)
}
