package main

import (
	"fmt"
)

func main()  {
	var high byte
	fmt.Println("请输入身高来判断体重")
	fmt.Scanln(&high)
	fmt.Println("体重为 = ",(high - 108)*2)
}