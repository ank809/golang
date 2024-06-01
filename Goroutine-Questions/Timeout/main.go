package main

import (
	"fmt"
	"time"
)

func sendValue(ch chan int) {
	for i := 0; i < 20; i++ {
		ch <- i + 1
		time.Sleep(time.Second * 4)
	}
}

func printValues(ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-time.After(time.Second * 2):
			fmt.Println("Time out")
		}

	}
}

func main() {
	ch := make(chan int)
	go printValues(ch)
	go sendValue(ch)
	time.Sleep(time.Second * 5)

}
