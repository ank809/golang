package main

import (
	"fmt"
	"time"
)

func sendOnly(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	close(ch)
}

func recieveOnly(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	go sendOnly(ch)
	go recieveOnly(ch)
	time.Sleep(time.Second * 2)

}
