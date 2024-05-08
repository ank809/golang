package main

import (
	"fmt"
	"time"
)

func SendValue(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func RecieveValue(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
func main() {
	ch := make(chan int)
	go SendValue(ch)
	time.Sleep(time.Second * 5)
	go RecieveValue(ch)
	time.Sleep(time.Second * 2)
}
