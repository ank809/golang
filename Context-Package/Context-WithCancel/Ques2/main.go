/*
	Ques

Context with a timeout of 10 seconds.
Start a goroutine to process the image.
The image processing involves:
Simulating a resize operation that takes 5 seconds.
Simulating a store operation that takes 3 seconds.
Each step should periodically check the context to see if it has been cancelled and terminate early if so.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	go imageProcessing(ctx)

	go func() {
		time.Sleep(time.Second * 7)
		cancel()
		fmt.Println("Context Cancelled")
	}()
	time.Sleep(time.Second * 9)
}

func imageProcessing(ctx context.Context) {
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("Image Successfully Resized")
	case <-ctx.Done():
		fmt.Println("Image not resized")
		return
	}

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("Image Successfully stored")
	case <-ctx.Done():
		fmt.Println("Image not stored")
		return
	}
}
