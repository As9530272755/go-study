package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Brethday string
	Sal      int
	Skill    string
}

// 将 json 字符串反序列化成结构体
func unmarshalStruct() {
	// 先拿到一个 json 的结构体字符串，然后用 \ 将双引号转义
	// str 在工作中是通过网络传输获取的，或者是读取文件获取到的
	str := "{\"name\":\"绿帽王\",\"age\":500,\"brethday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"牛魔拳\"}"

	// 定义一个 monster 实例
	var monster Monster

	// 调用 json 里面的函数
	// 这是 &monster 是引用传递，因为不用引用传递是改变不了 Monster 结构体的值
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Print("反序列化失败", err)
	}

	fmt.Print("反序列化后 monster.Name = ", monster.Name)

}

func main() {
	unmarshalStruct()
}
