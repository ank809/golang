package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)

	defer cancel()

	go printMessage(ctx)
	// time.Sleep(time.Second * 5)

	// main is waiting
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out")
	}

}

func printMessage(ctx context.Context) {
	select {
	case <-time.After(time.Second * 4):
		fmt.Println("Message sent successfully")
	case <-ctx.Done():
		fmt.Println("Task Time out")
	}
}
