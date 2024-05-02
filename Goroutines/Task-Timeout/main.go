package main

import (
	"fmt"
	"time"
)

func SendValues(ch chan int) {
	// time.Sleep(time.Second * 4)
	ch <- 10
}

func TimeOut(ch chan int) {
	select {
	case val := <-ch:
		fmt.Println(val)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout")
		// close(ch)
		return
	}
}
func main() {
	ch := make(chan int)
	go SendValues(ch)
	go TimeOut(ch)
	time.Sleep(time.Second * 10)

}
