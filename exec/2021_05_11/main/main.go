package main

import (
	"fmt"
)

// 定义 A 接口，定义了方法是 Say()
type A interface {

	// 而且我在这个 Say() 方法中定义的形参是 qiqi 类型为 string
	Say(qiqi string) string
}

// 定义一个 stu 结构体，字段为 name 类型为 string
type Stu struct {
	Name string
}

// 编写一个 Say() 方法，绑定给 Stu 结构体，从而实现了调用了 A 接口
func (stu *Stu) Say(name string) string {
	stu.Name = name
	str := fmt.Sprintf("name = %v", stu.Name)
	return str
}

func main() {
	stu := Stu{}
	var a A = &stu
	str := a.Say("黄二狗")
	fmt.Println(str)
}
