// package main

// import (
// 	"fmt"
// 	"time"
// )

// func PrintVal(ch chan int) {
// 	// for {
// 	fmt.Println(<-ch)
// 	// }
// 	close(ch)
// }

// func main() {
// 	ch := make(chan int)
// 	go PrintVal(ch)
// 	ch <- 12
// 	// This will cause a deadlock as we are sending on value on a closed channel
// 	ch <- 15
// 	time.Sleep(time.Second * 10)

// }

package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // go here if c is full.
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // go here if c is empty
		}
	}
	trySend("Hello!") // succeed to send
	// time.Sleep(time.Second * 2)
	trySend("Hi!") // succeed to send
	// Fail to send, but will not block.
	// time.Sleep(time.Second * 2)
	trySend("Bye!")
	// The following two lines will
	// both succeed to receive.
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// The following line fails to receive.
	fmt.Println(tryReceive()) // -
}
