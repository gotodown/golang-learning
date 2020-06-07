package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// fmt.Print()
	flagPa()
}

func flagPa() {
	name := flag.String("name", "liujiadong", "请输入本人的名字")
	age := flag.String("age", "30", "请输入本人的年龄")
	flag.Parse() // 使用flag

	fmt.Println(*name, *age)
}

// 获取命令行参数
// 直接读取命令行， 第一个参数是本身文件名
func fromArgs() {
	fmt.Printf("%#v\n", os.Args)
}
