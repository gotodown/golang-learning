package main

import "fmt"

// Action s
type Action struct {
	Legs int64
	Run  bool
}

// Animal 匿名结构体嵌套
// 匿名结构体嵌套在初始化时必须采用x.Action=Action{...} 或x.Action.Run的方式
// 但是调用可以直接使用x.Run, x.Legs的方式
type Animal struct {
	Kind string
	Name string
	Action
}

func main() {
	human := &Animal{
		Kind: "Monkey",
		Name: "Human Being",
	}
	human.Action = Action{
		Legs: 2,
		Run:  true,
	}
	fmt.Println(human.Run)
}
