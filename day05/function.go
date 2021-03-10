package main

import "fmt"

// Person Human being
type Person struct {
	Name    string
	Age     int
	Address string
}

// Shopping high
func (p Person) Shopping(money int) int {
	fmt.Println("花费了太多钱！")
	return money * 12
}

// GetName RETURN THIS PERSON'S NAME
func (p Person) GetName() string {
	return p.Name
}

// SetAge set this person's age
func (p Person) SetAge(age int) {
	p.Age = age
}

func main() {
	p := &Person{"liujiadong", 24, "shenzhen"}
	p.SetAge(23)
	fmt.Println(p.Age)
	fmt.Println(p.Shopping(1200))
	fmt.Println(p.GetName())

}
