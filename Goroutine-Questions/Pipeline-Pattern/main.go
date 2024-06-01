package main

import (
	"fmt"
	"time"
)

// Stage 1: to generate a series a numbers

func numbers(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i + 1
	}
	defer close(ch)
}

// Stage 2 : to square each number in channel

func squarenumbers(out chan<- int, in <-chan int) {
	for val := range in {
		out <- val * val
	}
	defer close(out)
}

// Stage 3: Print the values

func printValues(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go numbers(ch)
	go squarenumbers(ch2, ch)
	go printValues(ch2)
	time.Sleep(time.Second * 4)

}
