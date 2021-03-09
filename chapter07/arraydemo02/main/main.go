package main

import (
	"fmt"
)

func main() {
	// 使用数组的方式来解决
	
	// 1. 定义一个数组
	var hens [7]float64

	// 2.初始化数组给数组的每个元素赋值
	hens[0] = 3.0 // hens 数组的第一个元素赋值，下标是从 0 开始计算的，因为在计算机中 0 为 1
	hens[1] = 5.0 // 2个 依此类推
	hens[2] = 1.0 
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0
	hens[6] = 100.0

	// 3.数组可以遍历，我们遍历数组，求出总体重
	total := 0.0
	for i := 0 ; i < len(hens) ; i++ {
		total += hens[i]
	}

	// 4.求出平均体重
	average := fmt.Sprintf("%.2f",total / float64(len(hens)))
	fmt.Printf("总体重 = %v\n平均体重 = %v",total,average)
}