// Cancel context when user enters something

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		fmt.Println("Press Enter to cancel the context")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Context Cancelled", ctx.Err())
	case <-time.After(time.Second * 5):
		fmt.Println("Work completed")
	}
}
