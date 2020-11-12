package main

import (
	"fmt"
)

func main(){
	var price float32 = 89.12
	fmt.Println("price=",price)
	var num1 float32 = -0.000089
	var num2 float64 = -78123123.0123
	fmt.Println("num1=",num1,"num2=",num2)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=",num3,"num4",num4)

	var num5 = 1.1
	fmt.Printf("num5 的数据类型位 %T\n",num5)

	num6 := 5.12
	num7 := .123
	fmt.Println("num6=",num6,"num7=",num7)

	num8 := 5.1234e2
	fmt.Println("num8=",num8)

	num9 := 5.123E-2
	fmt.Println("num9=",num9)
}