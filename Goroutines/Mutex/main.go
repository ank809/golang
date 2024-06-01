package main

import (
	"fmt"
	"sync"
	"time"
)

var shared int = 10

var mutex = sync.Mutex{}

func Increment() {
	mutex.Lock()
	shared += 10
	mutex.Unlock()
}

func Read() {
	mutex.Lock()
	fmt.Print(shared)
	mutex.Unlock()
}

func main() {
	go Increment()
	go Read()
	fmt.Println("In main", shared)
	time.Sleep(time.Second)
}
