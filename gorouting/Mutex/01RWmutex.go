package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	lock   sync.Mutex
	wg     sync.WaitGroup
	rwlock sync.RWMutex
	once   sync.Once
)

func writer() {

	once.Do()
	lock.Lock() // 互斥锁
	x += 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10 ms
	lock.Unlock()
	// lock.Unlock()
	wg.Done()
}

func reader() {
	lock.Lock()
	time.Sleep(time.Millisecond) // 假设读操作耗时1ms
	lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writer()
	}
	for i := 0; i < 1000; i++ {
		if i%100 == 0 {
			fmt.Println("真的启动了！！！")
		}
		wg.Add(1)
		go reader()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(x)
}
