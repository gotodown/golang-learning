package main

import (
	"fmt"
	"time"
)

func main() {
	timeTest()
}

func timeTest() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.UnixNano())

	ret := time.Unix(1586594190217526622, 0)
	fmt.Println(ret)
	fmt.Println(time.Unix())
}
