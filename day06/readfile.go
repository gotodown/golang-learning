package main

import (
	"fmt"
	"os"
)

func main() {
	src := "/root/tmp/mysql/mysql-community-8.0.18-1.el7.src.rpm"
	tfile := src + "-tmp.txt" // 记录访问游标的位置
	ft, _ := os.OpenFile(tfile, os.O_CREATE|os.O_RDWR, 0644)
	tmp := make([]byte, 100, 100)
	n1, _ := ft.Read(tmp)
	fmt.Println("读取了。。。", n1)
}
