package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// MysqlConfig ...
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:“username”`
	Password string `ini:“password”`
}

// RedisConfig 配置字段 ...
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:“password”`
	Database int    `ini:”database“`
}

func loadConfig(filename string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	//data 的类型必须为结构体指针，
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		err := errors.New("data parmas must be a struct pointer ")
		return err
	}

	// 读文件得到字节数据
	in, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	context := string(in)
	lines := strings.Split(context, "\n")
	fmt.Println(lines)
	// 逐行读取
	var structName string
	for idx, line := range lines {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			// v := reflect.ValueOf(data)
			for i := 0; i < t.NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 确保配置为key=value 的格式
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return

			}
		}

	}
	return
}

func main() {
	var mc MysqlConfig
	loadConfig("./properties.ini", &mc)
}
