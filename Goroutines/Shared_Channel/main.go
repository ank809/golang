package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates random number, Prints random number

func GenerateRandom(ch chan int) {
	for i := 0; i < 10; i++ {
		randomnum := rand.Int()
		ch <- randomnum
	}
}

func ReadNumber(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
func main() {
	ch := make(chan int)

	go GenerateRandom(ch)
	go ReadNumber(ch)

	time.Sleep(time.Second)
}
