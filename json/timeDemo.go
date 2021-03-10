package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Post struct {
	CreatedAt time.Time `json:"created_at"`
}

func timeFieldDemo2() {
	p1 := Post{CreatedAt: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json")
		fmt.Printf("json.Unmarshal faile, err: %v\n", err)
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		fmt.Printf("JSON.")
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
func timeFieldDemo() {
	p1 := Post{CreatedAt: time.Now()}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("time is %s\n", b)
	t1 := `{"created_at":"2020-12-27T17:23:00.7497427+08:00"}`
	var p2 Post
	if err = json.Unmarshal([]byte(t1), &p2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("p2:%#v\n", p2)
}
func main() {
	timeFieldDemo2()
}
