package main

import "fmt"

/* 定义接口 */
type animal interface {
	SayHi()
	Run()
	CanFly() bool
}
type Person struct {
	Name string
}

func (p *Person) SayHi() {
	fmt.Println("Hello, I'm Person!!")
}

func (p *Person) Run() {
	fmt.Println("I use two legs to run, and I can run so fast!!")
}

func (p *Person) CanFly() bool {
	fmt.Println("I can't to fly, but i want to fly!")
	return false
}

func (p *Person) GetName() string {
	return p.Name
}
func main() {
	var a animal
	a = &Person{
		Name: "liujiadong",
	}
	b := a.CanFly()
	fmt.Printf("can fly? %b", b)
	a.Run()
}
