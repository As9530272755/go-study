package main

import (
	"fmt"
)

func main()  {
	// 定义月份的变量
	var month int32
	fmt.Println("输入月份(1-12)：")
	fmt.Scanln(&month)
	// 定义年龄变量
	var age int32
	fmt.Println("输入年龄：")
	fmt.Scanln(&age)
	// 定义一个最高票价 60 的变量
	var fares float32 = 60.0
	// 调用 month 变量判断是否 >=4 <= 10 为旺季
	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("老人三折 =  %v！！",fares / 3)
		} else if age >= 18 {
			fmt.Printf("成年人票价 = %v ！",fares)
		} else {
			fmt.Printf("儿童半价 = %v ！",fares / 2)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println("淡季成人票价40！")
		} else {
			fmt.Println("淡季老人和儿童票价20")
		}		
	} 
}