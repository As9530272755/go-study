package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func testStruct() {
	monster := Monster{
		Name:  "一拳超人",
		Age:   23,
		Skill: "认真一拳",
	}

	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Print("序列化失败 err = ", err)
	}

	fmt.Print(string(data))
}

func main() {
	testStruct()
}
