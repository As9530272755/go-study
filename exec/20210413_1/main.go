package main

import (
	"fmt"
)

func cal(n1 float64,n2 float64,operator string) float64 {
	res := 0.0
	switch operator {
		case "+":
			res = n1 + n2
		case "-":
			res = n1 - n2
		case "*":
			res = n1 * n2
		case "/":
			res = n1 / n2
		default:
			fmt.Println("请输入正确的运算符！")
	}
	return res
}

func main()  {
	var n1 float64
	fmt.Println("请输入 n1 的值：")
	fmt.Scanln(&n1)

	var n2 float64
	fmt.Println("请输入 n2 的值：")
	fmt.Scanln(&n2)

	var operator string
	fmt.Println("请输入 (+ - * /) 的值：")
	fmt.Scanln(&operator)

	res := cal(n1,n2,operator)
	fmt.Println(res)
}