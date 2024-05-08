package main

import (
	"fmt"
	"time"
)

func PrintMsg(ch1, ch2 chan string) {
	for {
		select {
		case <-ch1:
			fmt.Println(<-ch1)
		case <-ch2:
			time.Sleep(time.Millisecond * 2)
			fmt.Println("PONG CALLED", <-ch2)
		}
	}
}

func Ping(ch chan string) {
	for {
		ch <- "Ping"
	}
}
func Pong(ch chan string) {
	for {
		ch <- "Pong"
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go Ping(ch1)
	go Pong(ch2)
	go PrintMsg(ch1, ch2)
	time.Sleep(time.Second * 10)
}
