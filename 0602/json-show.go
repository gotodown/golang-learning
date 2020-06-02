package main

import (
	"encoding/json"
	"fmt"
)

// Person human type
type Person struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Weight float64 `json:"weight,omitempty"`
}

// Card type
type Card struct {
	ID    int64   `json:"id,float64"`   //添加 string tag
	Score float64 `json:"score,string"` // 添加 string tag
}

func main() {
	jsonStr1 := `{"id": "123","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarsha jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}
