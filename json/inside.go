package main

import (
	"encoding/json"
	"fmt"
)

// 结构体嵌套

// User 任务
type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Profile
}

// Profile 增强
type Profile struct {
	ID     int    `json:"id,string"`
	Gender string `json:"gender"`
}

func main() {
	Demo()
}

func Demo() {
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // float64
}

func IntDemo() {
	s := `{"name": "liujiadong", "address":"shenzhen", "id": "1", "gender":"male"}`
	var user User
	err := json.Unmarshal([]byte(s), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name, user.Gender, user.ID)
	fmt.Printf("%T\n", user.ID)
}
