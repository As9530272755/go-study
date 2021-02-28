package main

import (
	"fmt"
)

func main()  {
	var M byte
	var age int32
	fmt.Println("请输入月份：")
	fmt.Scanln(&M)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	if M >= 4 && M <= 10 {
		if age >= 18 && age <= 60{
			fmt.Println("60")
		} else if age > 60 {
			fmt.Println("票价为：", 60 / 3 )
		} else {
			fmt.Println("票价为：", 60 / 2 )
		}
	} else if age >= 18 && age <= 60 {
		fmt.Println("40")
	} else {
		fmt.Println("20")
	}
}