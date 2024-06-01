package main

import (
	"fmt"
	"time"
)

func sendValue1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	close(ch)
}
func sendValue2(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	close(ch)
}

func printValues(ch1, ch2 chan int) {
	for {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Println(val, "First channel")
			} else {
				ch1 = nil
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Println(val, "Second channel")
			} else {
				ch2 = nil
			}
		}
		if ch1 == nil && ch2 == nil {
			fmt.Println("Finished")
			break
		}

	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendValue1(ch1)
	go sendValue2(ch2)
	printValues(ch1, ch2)
	time.Sleep(time.Second * 2)

}
