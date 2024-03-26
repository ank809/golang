package main

import (
	"fmt"
	"time"
)

func selectStatement() {
	one := make(chan string)
	two := make(chan string)
	quit := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		one <- "Hey"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		two <- "Hello"
	}()
	for {
		select {
		case res1 := <-one:
			fmt.Println("Response from 1st channel is", res1)

		case res2 := <-two:
			fmt.Println("Response from 2nd channel is", res2)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: Exiting...")
			close(quit)
			return
		}
	}
}
