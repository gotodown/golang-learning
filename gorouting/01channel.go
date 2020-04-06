package main

import "fmt"

var b chan int
var c chan int

func main() {
	b = make(chan int)
	// go func() {
	// 	x := <-b
	// 	fmt.Println(x)
	// }()
	b <- 10
	fmt.Println(b)
	c = make(chan int, 16)
	c <- 10
	fmt.Println(c)
}
