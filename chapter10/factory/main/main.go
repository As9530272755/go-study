package main

import (
	"fmt"

	// 引入 model 包，也就是写入 model 的路径
	"go_code/chapter10/factory/model"
)

func main()  {
	// // 创建一个 student 变量实例，然后类型是 model 包中的 Student 结构体
	// stu := model.Student {
	// 	Name : "model",
	// 	Score : 77.1,
	// }

	// 当 model 包中 student 结构体首字母小写，我们可以通过工厂模式来解决
	// 这里 stu 直接调用 model 包的 NewStudent 函数传入 "newmodel",66.0 两个值
	stu := model.NewStudent("newmodel",66.0)

	fmt.Println(*stu)

	// 取出单个字段的值，stu.Name 字段是调用的 model.NewStudent 函数
	// stu.GetScore() 调用的是 GetScore()
	fmt.Println("name = ",stu.Name,"Score = ",stu.GetScore())
}