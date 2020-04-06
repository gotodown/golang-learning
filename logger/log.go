package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

//	不同等级的日志输出
const (
	TRACE = 1
	DEBUG = 2
	INFO  = 3
	WARN  = 4
	ERROR = 5
)

/*
	需求：
		将log 写入终端或指定文件中
		根据配置文件，将指定等级及以上的日志写入指定输出
*/

var lev int
var out io.Writer

//Init 日志模块初始化
func Init(level int, output string) {
	lev = level
	if output == "terminal" {
		out = os.Stdout
	} else {
		path := "./tmp.log"
		obj, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("open file failed, err: %v\n", err)
		}
		out = obj

	}
}

// Trace 等级的日志输出
func Trace(a ...interface{}) {
	if TRACE < lev {
		return
	}
	format := fmt.Sprintf("[%s]-[%s]:", time.Now(), "TRACE")
	fmt.Fprintln(out, format, a)
}

// Debug something
func Debug(a ...interface{}) {
	fmt.Println(DEBUG < lev, "---", lev, INFO)
	if DEBUG < lev {
		return
	}
	format := fmt.Sprintf("[%s]-[%s]:", time.Now(), "DEBUG")
	fmt.Fprintln(out, format, a)

}

// Info something
func Info(a ...interface{}) {
	if INFO < lev {
		return
	}
	format := fmt.Sprintf("[%s]-[%s]:", time.Now(), "INFO")
	fmt.Fprintln(out, format, a)

}

//Warn something
func Warn(a ...interface{}) {
	if WARN < lev {
		return
	}
	format := fmt.Sprintf("[%s]-[%s]:", time.Now(), "WARNING")
	fmt.Fprintln(out, format, a)

}

//Error something
func Error(a ...interface{}) {
	if ERROR < lev {
		return
	}
	format := fmt.Sprintf("[%s]-[%s]:", time.Now(), "ERROR")
	fmt.Fprintln(out, format, a)

}
func main() {
	Init(INFO, "terminal")
	Trace("this is trace")
	Debug("[STRING]", "liujiadong")
	Info("[STRING]", "liujiadong")
	Warn("[STRING]", "liujiadong")
	Error("[STRING]", "liujiadong")
}
