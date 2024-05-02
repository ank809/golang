package main

import (
	"fmt"
	"time"
)

func SendIntegers(ch chan int) {
	// defer func() {
	// 	fmt.Println("Values khtm ho gyi ..")
	// }()
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		ch <- i
	}
}

func SendStrings(ch chan string) {
	// defer func() {
	// 	fmt.Println("Values khtm ho gyi ..")
	// }()
	time.Sleep(time.Second * 2)
	strings := [20]string{
		"apple", "banana", "cherry", "date", "elderberry",
		"fig", "grape", "honeydew", "kiwi", "lemon",
		"mango", "nectarine", "orange", "papaya", "quince",
		"raspberry", "strawberry", "tangerine", "ugli fruit", "watermelon",
	}
	for i := 0; i < 20; i++ {
		ch <- strings[i]
	}
}

func PrintValues(ch1 chan int, ch2 chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
			// default:
			// 	fmt.Println("Channels khali h ")
		}
	}
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go PrintValues(ch1, ch2)
	go SendIntegers(ch1)
	go SendStrings(ch2)
	time.Sleep(time.Second * 10)

}
