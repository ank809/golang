package main

import (
	"context"
	"fmt"
	"time"
)

type Name string

func main() {
	parentctx := context.WithValue(context.Background(), Name("Ankita"), "1234")
	childctx := context.WithValue(parentctx, Name("Astha"), "5678")

	go func(ctx context.Context) {
		value := ctx.Value(Name("Astha"))
		fmt.Println(value)
	}(childctx)

	time.Sleep(time.Second)
}
