package main

import (
	"fmt"
)

func main()  {
	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	var age byte
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	if month >= 4 && month <= 10 {
		if age >= 18 && age <= 60 {
			fmt.Println("旺季票价为",60)
		} else if age > 60{
			fmt.Println("旺季票价为",60/3)
		} else {
		fmt.Println("旺季票价为",60/2)
		}
	} else if age >= 18 && age <= 60 {
		fmt.Println("淡季票价为",40)
	} else {
		fmt.Println("淡季票价为",20)
	}
}