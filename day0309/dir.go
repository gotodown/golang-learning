package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dirname := "D:/CodeHub/gocode/golang-learning/day0306"
	dir, _ := ioutil.ReadDir(dirname)
	os.FileInfo.Name()
	if len(dir) == 0 {
		fmt.Println(dirname + " is empty dir!")
	} else {
		fmt.Println(dir)
		fmt.Println(dirname + " is not empty dir!")
	}
}
