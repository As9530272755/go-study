package main

import (
	"encoding/json"
	"fmt"
)

func SliceJson() {
	// 定义一个 slice 的切片，该切片的类型是一个 []map[string]interface{}
	var slice []map[string]interface{}

	// 声明并且定义 m1 这个 map
	m1 := make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 12
	m1["address"] = "北京"

	// 通过 append 将其把 m1 追加到 slice 切片中
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"墨西哥", "洛杉矶"}

	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Print("失败 err = ", err)
	}

	fmt.Print("序列化后 slice = ", string(data))

}

func main() {
	SliceJson()
}
