package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	obj, err := os.Open("./test01.go")
	defer obj.Close()
	if err != nil {
		fmt.Println("file open error!", err)
	}
	// 读文件
	b := make([]byte, 128)
	_, err = obj.Read(b)
	if err != nil {
		fmt.Println("read err", err)
	} else {
		fmt.Println(string(b))
	}
	// bufio read
	reader := bufio.NewReader(obj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, err:%v", err)
		}
		fmt.Print(line)
	}
	ret, err := ioutil.ReadFile("./test01.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
