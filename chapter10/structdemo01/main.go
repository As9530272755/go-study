package main

import (
	"fmt"
)

// 张老太养了两只猫猫:一只名字叫小白,今年3岁,白色。还有一只叫小花,今年100岁,花色。
// 每只猫的属性，名字、年龄、花色

// 使用结构体 struct 演示
// 定义一个 cat 结构体，将猫的各个字段/属性信息，放入到 cat 结构体进行管理
// 将这只猫的 姓名、年龄、颜色 放入到 Cat 这个结构体中
type Cat struct {
	Name string
	Age int
	Color string
}
 
func main()  {
	// 创建一个 cat 的变量实例,数据类型就是上面我们定义的 Cat
	// 当我们去声明或者是创建一个结构体变量时，其实她的空间就已经分配了
	var cat1 Cat
	// cat1 引用 Cat 结构体的 Name ，并赋值为 小白,下面同理
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	fmt.Println("第一只猫的信息:")
	fmt.Println("姓名 = ",cat1.Name)
	fmt.Println("年龄 = ",cat1.Age)
	fmt.Println("花色 = ",cat1.Color)

	fmt.Println()
	// 定义第二只，创建一个 cat2 的变量类型为 Cat 结构体
	var cat2 Cat
	cat2.Name = "小花"
	cat2.Age = 100
	cat2.Color = "花色"
	fmt.Println("第二只猫的信息:")
	fmt.Println("姓名 = ",cat2.Name)
	fmt.Println("年龄 = ",cat2.Age)
	fmt.Println("花色 = ",cat2.Color)
}