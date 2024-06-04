package main

import (
	"context"
	"fmt"
)

func main() {
	parentCtx := context.WithValue(context.Background(), "parentId", "123")

	childCtx := context.WithValue(parentCtx, "childId", 456)

	printValues(childCtx)
}

func printValues(ctx context.Context) {
	value1 := ctx.Value("parentId")
	value2 := ctx.Value("childId")
	fmt.Printf("Parent's Id %v Child's Id %v \n", value1, value2)
}
