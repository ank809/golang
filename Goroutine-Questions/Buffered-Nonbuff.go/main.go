package main

import (
	"fmt"
	"time"
)

func sendValues1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
}

func sendValues2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
}

func recieveValues(ch1, ch2 chan int) {
	for {
		select {
		case c1 := <-ch1:
			fmt.Println("Value from chan1", c1)
		case c2 := <-ch2:
			fmt.Println("Value from chan2", c2)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 4)
	go sendValues1(ch1)
	go sendValues2(ch2)
	go recieveValues(ch1, ch2)
	time.Sleep(time.Second * 10)

}
