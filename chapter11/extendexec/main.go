package main

import (
	"fmt"
)

type A struct {
	Name string
	Age int
}

// 定义 C 结构体，其中有个字段为 int 数据类型
type C struct {
	A
	int
}

func main()  {
	// 演示匿名字段是基本数据类型的使用
	c := C{}
	c.Name = "Cccc"
	c.Age = 10
	c.int = 20
	fmt.Println(c)
}