package main

import (
	"fmt"
	"time"
)

// sender function
func PrintHi(ch chan string) {
	for {
		ch <- "Hi"
	}
}

// sender func
func PrintHello(ch chan string) {
	for {
		ch <- "Hello"
	}
}

// reciever func
func PrintMsg(ch1, ch2 chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
			// default:
			// 	fmt.Println("END....")
		}
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go PrintHi(ch1)
	go PrintHello(ch2)

	go PrintMsg(ch1, ch2)

	time.Sleep(time.Second * 10)

}
