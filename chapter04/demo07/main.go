package main

import (
	"fmt"
)

func main()  {
	// 要求：可以从控制台接收用户的信息，【用户姓名，年龄，薪水，是否通过考试】
	// 1、先申明需要的变量
	var name string		// 姓名
	var age byte		// 年龄
	var sal float32		// 薪水
	var isPass bool 	// 是否通过考试

	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)	// 实际上将 name 的变量信息传递给了 Scanln 函数，当程序执行到这个位置的时候会等待用户的输入并回车

	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)	// 将 age 变量参数值传递给 Scanln ，当程序执行到这个位置的时候会等待用户的输入并回车
	
	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)    // 将 sal 变量参数值传递给 Scanln ，当程序执行到这个位置的时候会等待用户的输入并回车

	// fmt.Println("请输入是否通过考试：")
	// fmt.Scanln(&isPass) // 将 isPass 变量参数值传递给 Scanln ，当程序执行到这个位置的时候会等待用户的输入并回车

	// fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n",name,age,sal,isPass)

	fmt.Println("请输入你的姓名，年龄，薪水，是否通过考试，输入时请使用空格隔开")
	fmt.Scanf("%s %d %f %t",&name ,&age ,&sal ,&isPass)
	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过考试 %v \n",name,age,sal,isPass)
}