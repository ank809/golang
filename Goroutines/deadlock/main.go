package main

import (
	"fmt"
	"time"
)

func Receiver(ch chan int) {
	val := <-ch
	fmt.Println(val)
}

func channel_num(number chan int) {

	number <- 239 //function with a data

}

func main() {
	num := make(chan int, 6) //initializing the channel

	go channel_num(num) //calling as a goroutine
	go Receiver(num)
	i := <-num //sending data from channel num
	//j := <-num //sending data from channel num

	/* Here deadlock occurs because we only send one value to the channel and we are trying to access
	more than one value i.e in i and j so this statement waits until the second value id recieved
	so deadlock occurs
	Only the "Value of Channel i,j =", i will execute not the one in Reciever func because it totally
	depends on the timing of the goroutine execution  If the Receiver goroutine executes before the
	main goroutine receives the value from the channel, it will print the value.
	*/
	fmt.Println("Value of Channel i,j =", i)
	time.Sleep(time.Second * 4)
}
