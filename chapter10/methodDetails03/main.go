package main

import (
	"fmt"
)

// 定义 student 结构体
type Student struct {
	Name string
	Age int
}

// 给他 student 实现 String() 方法
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}

func main()  {
	// 先定义一个 mainStu 变量,类型为 Student 结构体
	mainStu := Student{
		Name : "tom",
		Age : 20,
	}

	// 这里没有调用方法，而是直接通过 fmt.Println 默认会调用这个变量的 String() 方法进行输出
	// 这里通过 & 取地址符将 mainStu 传递给 String() 方法
	fmt.Println(&mainStu)
}