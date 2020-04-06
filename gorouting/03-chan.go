package main

import (
	"fmt"
	"sync"
)


func main() {
	jobs := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		// fmt.Println("GoStart")
		defer wg.Done()
		for i := 1; ; i++ {
			fmt.Println("GoforSTART", i)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
			// fmt.Println("GoforEND", i)
		}
	}()
	wg.Add(1)
	for j := 1; j <= 3; j++ {
		// fmt.Println("OutFOR", j)
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
