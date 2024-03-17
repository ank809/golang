// struct is a composite data type that groups together zero or more fields of possibly different data types.
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Rectangle struct {
	length, breadth int
}

func changeName(p *Person) {
	p.name = "Larry"
}

// Embedded Structs

type Shape struct {
	area int
	r    *Rectangle
}

func main() {
	var p1 Person = Person{"Ankita", 19}
	fmt.Println(p1.name, p1.age)

	p1.name = "Astha"
	fmt.Println(p1.name, p1.age)

	// r1 := Rectangle{10, 12}
	// fmt.Println(r1)

	r1 := &Rectangle{10, 12}
	fmt.Println(*r1)

	// If you only want to assign one value
	r2 := Rectangle{length: 12}
	fmt.Println(r2)
	changeName(&p1)
	fmt.Println(p1)

	s1 := Shape{12, &Rectangle{2, 4}}
	fmt.Println(s1)
	fmt.Println(s1.r)
	fmt.Println(s1.r.length)
	fmt.Println(s1.r.breadth)

	p3 := Point{12, 19}
	p4 := Point{9, 20}
	ans := distance(p3, p4)
	fmt.Println(ans)

	distance := p3.Distance(p4)
	fmt.Println(distance)

	student := Student{"Ankita", []int{80, 60, 90}, 19}

	fmt.Println(student.getAverage())

	fmt.Println(student.differnce_marks(Student{"Astha", []int{60, 40, 100}, 19}))

}
