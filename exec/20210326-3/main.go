package main

import (
	"fmt"
)

func main()  {
	// 定义一个 arr string 类型得数组
	var arr[10]string = [...]string{"AA","BB","CC","DD","AA","AA","EE","GG","MM","PP"}

	// 定义一个 find 初始值为 "AA" 的变量
	Find := "AA"
	for i := 0 ; i < len(arr) ; i++ {
		if Find == arr[i] {
		fmt.Printf("存在 %v 下标为 %v\n",Find,i)
		}
	}
	fmt.Println(arr)
}