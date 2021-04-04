package main

import (
	"fmt"
	// 调用 json 包
	"encoding/json"
)

// 定义 monster 结构体
type monster struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Skill string `json:"skill"`
}

func main()  {
	// 1.创建一个 monster 变量
	monster := monster{"牛魔王",100,"芭蕉扇"}

	// 2.将 monster 变量序列化为 json 格式的字符串
	// 使用的就是 json 包中的 Marshal 方法
	// 然后定义 JsonStr 变量来接收切片类型为 byte，err 变量来接收 Marshal 方法中的错误信息
	JsonStr,err := json.Marshal(monster)
	
	// 通过 if 判断 err 是否不等于 nil 空，如果不等于就说明有错误信息并将其输出
	if err != nil {
		fmt.Println("json 处理错误",err)
	}
	fmt.Println("JsonStr = ",string(JsonStr))
}