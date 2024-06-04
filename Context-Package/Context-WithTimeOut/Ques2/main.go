package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel1()
	defer cancel2()

	go printMessage(ctx1, "Task 1")
	go printMessage(ctx2, "Task 2")
	time.Sleep(time.Second * 4)
}

func printMessage(ctx context.Context, task string) {

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task timed out for ", task)
			return

		case <-time.After(time.Second):
			fmt.Println(task, "Running")
		}
	}

}
