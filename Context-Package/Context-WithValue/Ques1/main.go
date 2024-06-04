// Pass values from context to functions

package main

import (
	"context"
	"fmt"
)

type val struct {
	value string
}

func main() {
	v := val{
		value: "User Id",
	}
	ctx := context.WithValue(context.Background(), v, 123)
	// ctx := context.WithValue(context.Background(), "User Id", 123)

	printValues(ctx)
}

func printValues(ctx context.Context) {
	v := ctx.Value(val{value: "User Id"})
	fmt.Println(v)

	// fmt.Println(ctx.Value("User Id"))
}
