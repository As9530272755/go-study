package main

import (
	"encoding/json"
	"fmt"
)

// 编写
func test() string {
	var a []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "哈尼"
	m1["age"] = 12
	m1["address"] = "北京"

	m2 := make(map[string]interface{})
	m2["name"] = "老马"
	m2["age"] = 19
	m2["address"] = [2]string{"重庆", "北京"}

	a = append(a, m1)
	a = append(a, m2)
	data, err := json.Marshal(&a)
	if err != nil {
		fmt.Print("序列化 err = ", err)
	}

	return string(data)
}

// 将 json 字符串反序列化成 Map
func unmarshalSlice() {
	// 导入 json 数据
	// str := "[{\"address\":\"北京\",\"age\":12,\"name\":\"jack\"}," +
	//s	"{\"address\":[\"墨西哥\",\"洛杉矶\"],\"age\":20,\"name\":\"tom\"}]"
	str := test()

	// 定义一个 a 的 map 类型为 string 值为 interface{}
	var a []map[string]interface{}

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Print("反序列化失败", err)
	}

	fmt.Print(a)
}

func main() {
	unmarshalSlice()
}
