package main

import (
	"fmt"
	"time"
)

func Increment(ch chan int) {
	defer func() {
		fmt.Println("Increment closed")
		close(ch)
	}()

	for {
		val := <-ch
		if val == 6 {
			fmt.Println("Koi mil gya ")
			return
		}
		val += 1
		fmt.Println("Increment called", val)
	}
}

func Decrement(ch chan int) {
	defer func() {
		fmt.Println("Decrement closed")
		close(ch)
	}()

	for {
		val := <-ch
		if val == 6 {
			fmt.Println("Koi mil gya ")
			return
		}
		val -= 1

		fmt.Println("Decrement called", val)
	}
}

func Multiply(ch chan int) {
	defer func() {
		fmt.Println("Multiply closed")
		close(ch)
	}()

	for {
		val := <-ch
		if val == 6 {
			fmt.Println("Koi mil gya ")
			return
		}
		val *= 2
		fmt.Println("Multiply called", val)
	}
}

func main() {
	channels := [3]chan int{}

	for i := 0; i < 3; i++ {
		channels[i] = make(chan int)
	}
	go Increment(channels[0])
	go Decrement(channels[1])
	go Multiply(channels[2])

	for i := 0; i <= 6; i++ {
		for j := 0; j < 3; j++ {
			channels[j] <- i
		}
	}

	time.Sleep(time.Second * 10)
}
