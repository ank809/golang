package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomValue1(ch chan int) {
	for {
		val := rand.Intn(100)
		ch <- val
	}
}
func GenerateRandomValue2(ch chan int) {
	for {
		val := rand.Intn(100)
		ch <- val
	}
}
func GenerateRandomValue3(ch chan int) {
	for {
		val := rand.Intn(100)
		ch <- val
	}
}

func Reciever(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
func main() {
	ch := make(chan int)
	go Reciever(ch)
	go GenerateRandomValue1(ch)
	go GenerateRandomValue2(ch)
	go GenerateRandomValue3(ch)
	time.Sleep(time.Second * 10)

}
