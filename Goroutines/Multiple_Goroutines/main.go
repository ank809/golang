package main

import (
	"fmt"
	"time"
)

func SendValue1(ch chan int) {
	fmt.Println("1 called")
	for i := range 10 {
		ch <- i
	}
}
func SendValue2(ch chan int) {
	fmt.Println("2 called")
	for i := range 10 {
		ch <- i + 1
	}
}
func SendValue3(ch chan int) {
	fmt.Println("3 called")
	for i := range 10 {
		ch <- i + 2
	}
}
func SendValue4(ch chan int) {
	fmt.Println("4 called")
	for i := range 10 {
		ch <- i + 2
	}
}

func RecieveValue(ch chan int) {

	for {
		fmt.Println(<-ch)
	}
}
func main() {
	ch := make(chan int)
	go SendValue1(ch)
	go SendValue2(ch)
	go SendValue3(ch)
	go SendValue4(ch)
	go RecieveValue(ch)
	time.Sleep(time.Second * 5)
}
