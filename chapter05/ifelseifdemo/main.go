package main

import (
	"fmt"
)

func main()  {
	// var score int
	// fmt.Println("请输入岳小鹏成绩：")
	// fmt.Scanln(&score)
	// if score == 100 {
	// 	fmt.Println("奖励宝马")
	// } else if score >= 80 && score <= 99 {
	// 	fmt.Println("奖励一台iphone7plus")
	// } else if score >= 60 && score < 80{
	// 	fmt.Printf("奖励一个 iPad")	
	// } else { 
	// 	fmt.Printf("什么都没有")
	// }

	// 多分支使用陷阱
	var n int = 10
	if n > 9{
		fmt.Printf("ok1")
	} else if n > 6{
		fmt.Printf("ok2")
	} else if n > 3{
		fmt.Printf("ok3")
	} else {
		fmt.Printf("ok4")
	}
}