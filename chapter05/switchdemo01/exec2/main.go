package main

import (
	"fmt"
)

func main()  {
	var score float64
	fmt.Println("请输入对应的成绩")
	fmt.Scanln(&score)
	switch int(score / 60) {	// 这里我将 score 也就是我们输入的值和 60 做除法并且做 int 类型
		case 1:					// 除了之后等于 1 我们就输出合格
			fmt.Println("合格")
		case 0:					// 除了之后不等于 1 就不合格
			fmt.Println("不合格")
		default :
			fmt.Println("输入有误")
	}
}