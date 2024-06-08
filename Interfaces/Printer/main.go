package main

import "fmt"

type Printer interface {
	Print()
}
type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type AllInOnePrinter struct{}

func (a AllInOnePrinter) Print() {
	fmt.Println("Printer is printing")
}

func (a AllInOnePrinter) Scan() {
	fmt.Println("Scanner is scanning")
}

func main() {
	a := AllInOnePrinter{}
	a.Print()
	a.Scan()
}
