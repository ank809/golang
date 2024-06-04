package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*100))
	ch := make(chan string)
	defer cancel()

	url := "https://github.com/"

	go search(ctx, url, ch)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err().Error())
	case val := <-ch:
		fmt.Println(val)
	}

}

func search(ctx context.Context, url string, ch chan string) {
	req, err := http.NewRequestWithContext(ctx, "Get", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
	ch <- "Response Success"

}
