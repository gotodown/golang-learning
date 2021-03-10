package main

import "fmt"

type student struct {
	name string
	age int
}

func testSruct(){
	m := make(map[string]*student)
	stus := []student{
		{name: "ljy", age:30},
		{name: "liujiadong", age: 27},
		{name: "dx", age: 26},
	}

	for _, s := range stus{
		//fmt.Println(s.name)
		m[s.name] = &s
		fmt.Printf("%s->%p\n",s.name, &s)
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "==>", v.name)
	}
}
func (s *student)SetAge(age int){
	s.age = age
}
func main() {
	s := &student{
		name: "liujiadong",
		age: 30,
	}
	s.SetAge(20)
	fmt.Println(s.age)
}
