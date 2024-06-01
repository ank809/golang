package main

import (
	"fmt"
	"time"
)

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	close(ch)
}

func recieve(ch chan int) {
	// for val := range ch {
	for {
		// fmt.Println(val)
		// fmt.Println(<-ch)
		val, open := <-ch
		if !open {
			break
		}
		fmt.Println(val, open)
	}
	// close(ch)
}

func main() {
	// send only
	ch1 := make(chan int)
	go send(ch1)
	go recieve(ch1)
	time.Sleep(time.Second * 2)
}
