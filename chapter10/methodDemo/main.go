package main

import (
	"fmt"
)

// 定义一个结构体 Person 字段为 Num
type Person struct {
	Name string
}

// 相当于给 person 这个类型绑定一个方法，这个方法也就是输出 person 结构体的 Name
// a 相当于是一个结构体 person 的变量名
// test() 是这个方法的名字
func (a Person) test() {
	a.Name = "方法"
	// 输出 a.Num 相当于输出 Person.Num
	fmt.Println("test()",a.Name)
}

func main() {
	// 定义 p 变量类型为 person
	var p Person
	
	// 给 p.Name 赋值 zhang
	// 一旦 p.Name 有值了，然后在传给 test 方法
	p.Name = "zhang"

	// 调用 test 方法
	// 这个 p 会在调用 test 方法的时候传给 test 方法中的 a
	p.test()
	fmt.Println("main()",p.Name)
}