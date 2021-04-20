package main

import (
	"go_code/MyAccount/utils"
)

func main()  {
	// 返回值是一个结构体，然后调用了 MainMenu 这个公开的方法
	MyAccount := utils.NewmyAccount()
	MyAccount.MainMenu()
}