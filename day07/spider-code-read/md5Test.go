package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// test := map[[md5.Size]byte]struct{}{}
	wg.Add(1)
	go func() {
		fmt.Println("something is here")
		wg.Done()
	}()
	fmt.Println(SprintStack())
	wg.Wait()

}

// SprintStack 打印
func SprintStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], true)
	return string(buf[:n])
}
