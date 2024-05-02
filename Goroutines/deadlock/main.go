package main

import (
	"fmt"
)

// func Receiver(ch chan int) {
// 	val := <-ch
// 	fmt.Println(val)
// }

func channel_num(number chan int) {

	number <- 239 //function with a data

}

func main() {
	num := make(chan int, 6) //initializing the channel

	go channel_num(num) //calling as a goroutine

	i := <-num //sending data from channel num
	j := <-num //sending data from channel num
	fmt.Println("Value of Channel i,j =", i, j)
}
