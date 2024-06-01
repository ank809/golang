package main

import (
	"fmt"
	"math"
	"time"
)

func PrimeNumberGenerator(ch chan int) {
	defer close(ch)
	for num := 1; num < 100; num++ {
		count := 0
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				count++
			}

		}
		if count == 0 {
			ch <- num
		}
	}
}

func DisplayNumbers(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func main() {
	ch := make(chan int)
	// DisplayNumbers(ch)
	go PrimeNumberGenerator(ch)
	go DisplayNumbers(ch)
	time.Sleep(time.Second * 5)
}
