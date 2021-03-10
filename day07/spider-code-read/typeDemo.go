// 测试类型转换
package main

import "fmt"

// Demo test
type Demo int64

func main() {
	s := intT()
	Use(s)
}

func intT() func(i int64) {
	return func(i int64) {
		fmt.Println("这是一个匿名函数", i)
	}
}

// Use asd
func Use(e ...interface{}) {
	a := int64(64)
	for _, fn := range e {
		switch fn.(type) {
		case func(b int64):
			fn.(func(b int64))(b)
		default:
			return
		}
	}
}
