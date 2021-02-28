package main

import (
	"fmt"
)

func main()  {
	var second float64
	fmt.Println("请输入比赛成绩(秒)：")
	fmt.Scanln(&second)
	if second <= 8 {
		fmt.Println("进入决赛")
		
		var gender string
		fmt.Println("请输入性别(男/女)：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入男子组")
		} else if gender == "女" {
			fmt.Println("进入女子组")
		}
	} else {
		fmt.Println("淘汰！")
	}
}