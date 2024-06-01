package main

// import (
// 	"fmt"
// 	"time"
// )

// func sendValue(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i + 1
// 	}
// 	close(ch)
// }

// func recieveValues(ch chan int) {
// 	for val := range ch {
// 		fmt.Println(val)

// 	}
// }
// func main() {
// 	ch := make(chan int)
// 	go sendValue(ch)
// 	go recieveValues(ch)
// 	time.Sleep(time.Second * 5)package main

import (
	"fmt"
)

func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Close the channel after sending all values
}

func main() {
	ch := make(chan int)

	go sendData(ch)

	// Receive values from the channel until it's closed
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, exiting loop.")
			break
		}
		fmt.Println("Received:", val)
	}

	// Attempt to receive from the closed channel again
	val, ok := <-ch
	if !ok {
		fmt.Println("Tried to receive from closed channel, ok value:", ok)
	} else {
		fmt.Println("Received from closed channel:", val)
	}
}

// }
