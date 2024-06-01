package main

import (
	"fmt"
	"time"
)

func sendValues(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func printValues(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int, 4)
	// go printValues(ch)
	go sendValues(ch)
	go printValues(ch)
	time.Sleep(time.Second)
}

// import (
// 	"fmt"
// 	"time"
// )

// func sender(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Sending:", i)
// 		ch <- i
// 	}
// 	close(ch)
// }

// func main() {
// 	// Non-buffered channel
// 	ch1 := make(chan int)
// 	go sender(ch1)

// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Receiving from non-buffered channel:")
// 	for v := range ch1 {
// 		fmt.Println("Received:", v)
// 	}

// 	// Buffered channel
// 	ch2 := make(chan int, 3)
// 	go sender(ch2)
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("Receiving from buffered channel:")
// 	for v := range ch2 {
// 		fmt.Println("Received:", v)
// 	}
// }
