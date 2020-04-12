package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup

// var lock sync.Mutex

func add() {
	for i := 0; i < 100; i++ {
		// lock.Lock()
		x += i
		// lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(3)
	go add()
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
