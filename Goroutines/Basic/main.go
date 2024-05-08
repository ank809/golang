package main

// We can send the value inside a channel even if it is closed if the capacity is not full but if
// it is then a deadlock will occur as value send on a closed channel
import (
	"fmt"
	"time"
)

func waiting(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
	// close(ch)
}

func main() {
	ch := make(chan int, 3)
	go waiting(ch)
	ch <- 12
	ch <- 15
	time.Sleep(time.Second * 2)
}
