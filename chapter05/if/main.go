package main

import (
	"fmt"
)

func main()  {
	// 编写一个程序，可以输入人的年龄，如果该同志的年龄大于18岁，则输出“你年龄大于18，要对自己的行为负责！”

	// 分析：
	// 1.使用一个变量来表示年龄 var age int
	// 2.冲控制台接收一个输入 fmt.Scanln(&age) 函数
	// 3.if 判断这个输入的数字是否大于 18

	var age int
	fmt.Println("输出一个年龄：")
	fmt.Scanln(&age)
	if  age > 18 {
		println ("你年龄大于18，要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不大这次放过你了")
	}
}