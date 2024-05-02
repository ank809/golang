// package main

// import (
// 	"fmt"
// 	// "sync"
// )

// func buyTicket(remaining_tickets *int, id int) {
// 	if *remaining_tickets > 0 {
// 		*remaining_tickets--
// 		fmt.Printf("Ticket bought by user %d and Remaining tickets are %d \n", id, *remaining_tickets)
// 	} else {
// 		fmt.Printf("No tickets left for user %d\n", id)
// 	}
// 	// defer wg.Done()
// }

// func main() {

// 	// var wg sync.WaitGroup

// 	tickets := 500

// 	for i := 0; i < 2000; i++ {
// 		// wg.Add(1)
// 		go buyTicket(&tickets, i)
// 	}
// }

package main

import "fmt"

func main() {
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i+x)
		}
		x = 10
	}()
	fmt.Println()
	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i+x)
			}()
		}
		x = 10
	}()
}
