package main

import (
	"fmt"
)

//定义一个数组，并给出 8 个整数，求该数组中大于平均值的数的个数，和小于平均值的数的个数。

func main()  {
	// 定义数组
	arr := [8]int{10,20,30,40,30,60,70,80}

	// 定义 sum 变量，用于取出总数
	sum := 0
	for i :=0 ; i < len(arr) ; i++ {
		sum += arr[i]
	}

	// 定义 ave 变量，用于取出平均值
	ave := sum / len(arr)

	// 定义 max ，用于记录大于 ave 平均数的个数
	max := 0

	// 定义 small ，用于记录小于 ave 平均数的个数
	small := 0

	for i := 0 ; i < len(arr) ; i++ {
		// 如果 ave 平均数小于 arr[i] 遍历的数组
		if ave < arr[i] {
			
			// 大于平均数max 变量就依次累计
			max++
		// 如果 ave 大于 arr[i] 遍历的数组
		} else if ave > arr[i] {
			// 则小于平均数 small 依次累计
			small++
		}
	}
	fmt.Println("数组是",arr)
	fmt.Println("总数：",sum)
	fmt.Println("平均数：",ave)
	fmt.Println("大于平均数的个数",max)
	fmt.Println("小于平均数的个数",small)
}