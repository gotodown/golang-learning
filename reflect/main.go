package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	// reflect.Method.Name
	fmt.Println(t)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is flaot32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func test() {
	fmt.Println(123)
}
func main() {

	stu := student{
		Name:  "liujiadong",
		Score: 81,
	}
	t := reflect.TypeOf(stu)
	fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		feild := t.Field(i)
		fmt.Printf("name:%s, index:%d, type:%v json tag :%v \n", feild.Name, feild.Index, feild.Type, feild.Tag)
		fmt.Println(feild)
	}
	// var a int64 = 3
	// reflectValue(a)
}
