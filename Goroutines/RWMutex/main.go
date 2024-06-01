package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.RWMutex{}
var shared int = 10

func ReadValue1() {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("Read 1", shared)

}
func ReadValue2() {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("Read 2", shared)
}
func ReadValue3() {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("Read 3", shared)
}
func ReadValue4() {
	mutex.RLock()
	defer mutex.RUnlock()
	fmt.Println("Read 4", shared)
}

func WriteValue() {
	mutex.Lock()
	shared += 1
	mutex.Unlock()
}

func main() {
	go WriteValue()
	go ReadValue1()
	go ReadValue2()
	go ReadValue3()
	go ReadValue4()
	time.Sleep(time.Second)
}
