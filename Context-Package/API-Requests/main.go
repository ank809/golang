package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// All the API requests are canceled if any of them exceeds a specified timeout.

func main() {
	urls := []string{
		"https://www.google.com/",
		"https://www.youtube.com/",
		"https://twitter.com/home",
		"https://www.instagram.com/",
	}

	ch := make(chan string)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, url := range urls {
		go fetchUrl(ctx, url, ch)
	}

	for res := range ch {
		fmt.Println(res)
	}
}

func fetchUrl(ctx context.Context, url string, result chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "Get", url, nil)
	if err != nil {
		fmt.Println(err)
		// Creates formatted string without printing it directly to the console
		result <- fmt.Sprintf("Error creating request for %s ", url)
		return
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	defer resp.Body.Close()
	result <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)

}
