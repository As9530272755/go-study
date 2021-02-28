package utils

import (
	"fmt"
)

	var Num1 int = 10

	// 为了让其它包的文件使用 Cal 函数，需要将 C 大写类似其它编程语言的 public
func Cal(n1 float64,n2 float64,operator byte) (float64)  {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 /n2
	default :
		fmt.Println("请输入正确的运算符号")
	}
	return res
}  