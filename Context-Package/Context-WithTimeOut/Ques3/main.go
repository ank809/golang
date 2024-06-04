/* parent context with a 5-second timeout and a child context with a 2-second timeout. Start a goroutine that simulates a 3-second task. */

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	parentCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	childCtx, cancelChild := context.WithTimeout(parentCtx, time.Second*2)

	defer cancelChild()

	go printMessage(parentCtx, "Parent Context")
	go printMessage(childCtx, "Child Context")
	time.Sleep(time.Second * 6)
}

func printMessage(ctx context.Context, val string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context timed out ", val)
			return
		case <-time.After(time.Second * 3):
			fmt.Println(val)
		}

	}
}
