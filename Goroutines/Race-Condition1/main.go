package main

import (
	"fmt"
	"time"
)

var shared int = 12

func IncreaseVariable() {
	shared = shared + 10
}
func ReadVariable() {
	fmt.Println("Value of shared is ", shared)
}

func main() {
	go IncreaseVariable()
	time.Sleep(time.Second)
	go ReadVariable()
	time.Sleep(time.Second * 2)

}
