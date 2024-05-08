package main

import (
	"fmt"
	"time"
)

func PassValue(ch1 chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Pass Value Called", i+1)
		ch1 <- i + 1
	}
}

func PassValueToChan(ch1, ch2 chan int) {
	for val := range ch1 {
		fmt.Println("PassValueToChan Value Called", val)
		res := val + 4
		ch2 <- res
	}
}
func PrintVal(ch chan int) {
	for val := range ch {
		fmt.Println("Print Val", val)
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go PassValue(ch1)
	go PassValueToChan(ch1, ch2)
	go PrintVal(ch2)
	time.Sleep(time.Second * 5)

}
