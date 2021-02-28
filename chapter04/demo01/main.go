package main

import (
	"fmt"
)

func main()  {
	
	// 除法的案例
	
	// 说明，如果运算的数都是整数，那么除后去掉小数部分保留整数部分，也就说 10 / 4 = 2 
	fmt.Println(10 / 4)	

	var n1 float32 = 10 / 4 // 这里采用的是浮点数数据类型
	fmt.Println("n1=",n1)

	//如果我们希望保留小数部分，则需要浮点数参与运算，也就是将被除的数写成带小数类型
	var n2 float32 = 10.0 /4
	fmt.Println("n2=",n2)

	// 演示 % 取模的范例
	fmt.Println(10 % 3) 

	// 看一个公式 a % b = a - a / b * b

	fmt.Println("-10 %3 = ",-10 % 3) // = -10 - (-10) / 3 * 3 = -10 -3 * 3 = -10 - -9 = -1

	fmt.Println("10 % -3=",10 % -3) // = 10 - 10 / -3 * -3 = 10 - 3 * -3 = 10 - -9 = 1

	fmt.Println("-10 % -3 = ",-10 % -3) // -10 - (-10) / -3 * -3 = -10 - 3 * -3 = -10 - -9 = -1

	// ++ 和 -- 的使用
	var i int = 10
	i++
	fmt.Println("i=",i)

	var j int = 90
	j--
	fmt.Println("j=",j)
}