package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled")
				return
			case <-time.After(time.Second):
				fmt.Println("Task running ")
			}
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	cancel()

	time.Sleep(time.Second)
}
