package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLinkRepeating(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Maybe  down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link

}
