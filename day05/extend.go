/*
  继承
*/
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	human  Human //匿名字段
	school string
}

// Employee 继承
type Employee struct {
	Human   //匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// SayHi 重写
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s's %s you can call me on %s\n", e.company, e.name, e.phone)
}
func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.human.SayHi()
	sam.SayHi()
}
