package main

import (
	"fmt"
)

func main()  {
	
	// 求出一个数组的和还有平均值
	// 1.声明数组
	// 2.求出和 sum
	// 3.求出平均值
	var intArr[5]int = [...]int{20,21,20,20,20}
	
	// 定义 sum 变量来接收总和
	sum := 0
	for _,value := range intArr {
		sum += value
	}
	
	// 通过下面的方式将平均值保留到小数点位
	fmt.Printf("sum = %v\nave = %v",sum,float64(sum) / float64(len(intArr)))
}