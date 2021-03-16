package main

import (
	"fmt"
)

func main()  {
	// 使用常规的 for 循环遍历
	
	// 定义 arr 数组
	var arr [5]int = [...]int{10,20,30,40,50}

	// 使用切片来引入 arr 数组的 1-4 元素 20,30,40
	slice := arr[1:4]
	
	// 通过 for 循环来进行常规的遍历
	for i := 0 ; i < len(slice) ; i++{
		fmt.Printf("slice[%v] 遍历之后 = %v\n",i,slice[i])
	}

	// 使用 for-range 方式遍历
	var arr1 [5]int = [...]int{11,22,33,44,55}
	
	// 这里取出了 arr1 数组所有的元素
	slice1 := arr1[:]
	for index,value := range slice1 {
		fmt.Printf("slice1[%v] for-range = %v\n",index,value)
	}
}