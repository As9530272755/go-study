package main

import "fmt"

func main()  {
	var Name string
	var S int

	C := 3
	for i := 1 ; i <= 3 ; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&Name)
		if Name == "张无忌" {
			fmt.Println("请输入密码")
			fmt.Scanln(&S)
			if S == 888 {
				fmt.Println("登录成功")
				break
			}
		} else {
			fmt.Printf("还有 %v 次登录机会",C - i)
		}
		fmt.Println()
	} 
}