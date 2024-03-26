package main

import (
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://instagram.com",
		"http://github.com",
		// "htp://github.com",
	}
	// It'll return error only when the url you have passed is not in  valid format("http://")

	c := make(chan string)

	for _, link := range links {
		// checkLink(link)
		// go checkLinkisRunning(link, c)
		go checkLinkRepeating(link, c)

	}
	// Recieving messages from channel is a blocking thing
	// fmt.Println(<-c)

	// Receive messages from the channel for each link

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLinkRepeating(<-c, c)
	// }

	for l := range c {
		// go checkLinkRepeating(l, c)
		// function Literal
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkLinkRepeating(l, c)
		// }()
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkRepeating(link, c)
		}(l)
	}

	selectStatement()
}
