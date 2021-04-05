package main

import (
	"fmt"
)

type Calcuator struct {
	Num1 float64
	Num2 float64
}

/*
定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
实现形式2: 用一个方法搞定
*/

func (ma *Calcuator) Cal(ae byte) float64 {
	res := 0.0
	switch ae {
		case '+':
			res = ma.Num1 + ma.Num2
		case '-':
			res = ma.Num1 - ma.Num2
		case '*':
			res = ma.Num1 * ma.Num2
		case '/':
			res = ma.Num1 / ma.Num2
		default :
			fmt.Println("请输入正确的运算符")
	}
	return res
}


func main()  {
	var mu Calcuator
	mu.Num1 = 10
	mu.Num2 = 20
	res := mu.Cal('*')
	fmt.Println("结果：",res)
}