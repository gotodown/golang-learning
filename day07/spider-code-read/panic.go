package main

import "fmt"

func PanicT() {
	fmt.Println("我是大英雄！！")
	panic("oooooooooooooooooh")
}
func continues() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉到你了， bad g")
			fmt.Println(" wo haishi xiang rangta 继续执行")
		}
	}()
	PanicT()

}

func main() {
	continues()
	fmt.Println("wojixula")
}
