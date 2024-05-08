package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"

	fmt.Printf("ch %v len %v cap %v \n", ch, len(ch), cap(ch))
	close(ch)

	fmt.Printf("val from ch %v len %v cap %v \n", <-ch, len(ch), cap(ch))
	// This will print an empty string
	fmt.Printf("val from ch %v len %v cap %v \n", <-ch, len(ch), cap(ch))
	for i := range 10 {
		fmt.Printf("%v time , val from ch %v len %v cap %v \n", i+1, <-ch, len(ch), cap(ch))
	}

	// Will give an error as value send on a closed channel
	// ch <- "Hi"
	// fmt.Printf("val from ch %v len %v cap %v \n", <-ch, len(ch), cap(ch))

	// How to differentiate that the value you send on a channel is empty or it is the default one

	// False will denote that the val is default one and true will that it is send on channel
	// ch <- ""
	// val, ok := <-ch
	// fmt.Printf("%v val from channel %v \n", val, ok)
}
