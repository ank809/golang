package main

import (
	"fmt"
)

/*
To run multiple processes concurrently we have threads in java and in go we have goroutines and channels.
Concurrency is about managing multiple tasks and  without necessarily requiring them to run simultaneously.
Parallelism is about running multiple tasks at same time typically to improve the performance.
Concurrency do not guarantee that task will run in true paralel, they may share the resources and take turns to use them.
Parallelism ensures that all the task will execute at same time on separate resources.

Goroutines are lightweight independently executes functions that can run concurrently. Similar to threads but manages by go's runtime.
Channels are communication mechanism in go that allows go routines to communicate and synchronize.

TYPES OF CHANNELS :
1. Buffered Channels: They have a capacity to store  values and do not need reciever to recieves the value simultaneously until the capacity is full.Asynchronous
2. Unbuffered Channels :  They do not have any capacity and needs reciever to recieves the value otherwise it will become a block statement. */

func main() {
	// Nothing will be printed beacuse the main function terminates before the routine
	// So we need to make sure that our main function waits until the goroutines are terminated.
	// go printHello()
	ch := make(chan string)
	go printHello(ch)
	// time.Sleep(time.Millisecond * 5)
	fmt.Println(<-ch)

	chn := make(chan int)
	go printNumbers(chn)
	for i := 0; i < 8; i++ {
		fmt.Println(<-chn)
	}

}

func printHello(ch chan string) {
	// we can only pass one value at a time in channel in golang
	// ch <- "Hello"
	ch <- "Hey"
	close(ch)

}

func printNumbers(chn chan int) {
	for i := 0; i < 10; i++ {
		chn <- i

	}
	close(chn)
}
