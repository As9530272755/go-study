package main

import (
	"fmt"
)
/*
1) 编写一个 Student 结构体，包含 name、gender、age、id、score 字段，分别为 string、string、int、int、 float64 类型。
2) 结构体中声明一个 say 方法，返回值 string 类型，方法返回信息中包含所有字段值。
3) 在 main 方法中，创建 student 结构体实例(变量)，并访问say方法，并将调用结果打印输出。
*/

// 编写一个结构体，并存放要求字段
type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

// 定义一个 Say() 方法
func (stu *Student) Say() string {
	// 定义一个 infoStr 变量接收等会从 main 函数中传递过来的值
	infoStr := fmt.Sprintf("student 的信息 name = [%v] Gender = [%v] Age = [%v] Id = [%v] Score = [%v]",
	stu.Name,stu.Gender,stu.Age,stu.Id,stu.Score)
	
	// 返回值为 infoStr
	return infoStr
}

func main()  {
	// 测试
	// 创建要给 stu 变量，类型为 Student 结构体，并且为 Student 结构体的每个字段赋值
	var stu Student
	stu.Name = "tom"
	stu.Gender = "男"
	stu.Age = 29
	stu.Id = 1
	stu.Score = 77.7
	
	// 定义 str 变量接收 Say() 方法的返回值
	str := stu.Say()
	// 输出 str 变量
	fmt.Println(str)
}