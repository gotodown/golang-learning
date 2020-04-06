package main

import "fmt"

func main() {
	ch1 := make(chan int, 16)
	ch2 := make(chan int, 16)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			ch2 <- v
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}
