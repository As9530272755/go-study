package main

import (
	"encoding/json"
	"fmt"
)

func MapJson() {
	a := make(map[string]interface{})
	a["name"] = "孙悟空"
	a["age"] = 1000
	a["skill"] = "金箍棒"
	a["address"] = "水帘洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Print("序列化失败 err = ", err)
	}

	fmt.Print("序列化后 a = ", string(data))
}

func main() {
	MapJson()
}
