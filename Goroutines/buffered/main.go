package main

import (
	"fmt"
	"time"
)

func Receiver(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		ch <- i
	}

	time.Sleep(time.Second * 2)
	go Receiver(ch)
	time.Sleep(time.Second * 2)

}
