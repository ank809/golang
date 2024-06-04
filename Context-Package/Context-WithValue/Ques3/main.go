package main

import (
	"context"
	"fmt"
	"time"
)

type Val string

func main() {
	val := "1234"
	ctx := context.WithValue(context.Background(), Val(val), "Hello")
	ctx = context.WithValue(ctx, Val("Hi"), "234")

	go func(ctx context.Context) {
		val := ctx.Value(Val("1234"))
		fmt.Println(val)
	}(ctx)

	time.Sleep(time.Second * 2)
}
