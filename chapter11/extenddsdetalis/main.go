package main

import (
	"fmt"
)

// 定义 A 结构体，首字母大写的 Name 和首字母小写的 age
type A struct {
	Name string
	age int
}

// 给 A 结构体绑定 SayOk 方法首字母大写
func (a *A) SayOk() {
	fmt.Println("A 绑定 SayOk",a.Name)
}

// 给 A 结构体绑定 hello 方法首字母小写
func (a *A) hello() {
	fmt.Println("A 绑定 hello",a.Name)
}
// 可以看到 A 结构体中分别有两个字段其首字母分别是大小写、方法也是其首字母大小写


// 定义结构体 B，字段为 A 
type B struct {
	A
	// 这时候在 B 结构体中定义 Name
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B 绑定 sayok",b.Name)
}

func main()  {
	b := B{}

	// 简化后赋值
	b.Name = "b中NAME"
	b.age = 11
	b.A.Name = "A中Name"
	// 简化后调用 A 结构体中的方法
	b.SayOk()
	b.hello()
}