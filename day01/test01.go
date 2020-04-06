package main

import "fmt"

func main() {
	var s string
	s = "liujiadong"
	assign(s)
	a := 123
	assign(a)
}

func assign(a interface{}) {
	str, ok := a.(string)
	if !ok {
		fmt.Printf("param a is not string, it's %T\n", a)
	} else {
		fmt.Println("param a is string, thanks! a is ", str)
	}
}
