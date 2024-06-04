package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))

	defer cancel()

	deadline, ok := ctx.Deadline()

	if ok {
		fmt.Println("Context has deadline", deadline)
	} else {
		fmt.Println("Context do not have deadline")
	}

	select {
	case <-ctx.Done():
		fmt.Println("Cancelled because", ctx.Err())
	case <-time.After(time.Second * 3):
		fmt.Println("Work completed", time.Now())
	}
}
