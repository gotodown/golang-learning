package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全
var m = sync.Map{}

func get(key string) interface{} {
	val, _ := m.Load(key)
	return val
}
func set(key string, value int) {
	m.Store(key, value)
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k:=%v, v:= %v", key, get(key))
			fmt.Println()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
