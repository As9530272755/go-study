package main

import (
	"fmt"
)

// 编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数和对应的数组下标是多少?

// 编写一个接收函数的数组
func Arr(arr *[6]int)  {

	// 定义 MaxVla 变量和 Maxindex 变量来获取他的最大值和最大下标
	MaxVla := arr[0]
	Maxindex := 0

	// 定义 SmallVal 和 SmallIndex 变量来获取他的最小值和最小下标
	SmallVal := arr[0]
	SmallIndex := 0

	for i :=0 ;i < len(arr) ;i++ {
		// 判断 MaxVla 是否比 arr[i] 小，如果比 arr[i] 小就 MaxVla = arr[i],如果 MaxVla 大于 arr[i] 就不赋值
		if MaxVla < arr[i] {
			MaxVla = arr[i]
			Maxindex = i
		}
		// 判断 SmallVal 是否比 arr[i] 大，如果比 arr[i] 大就 SmallVal = arr[i],如果 SmallVal 小于 arr[i] 就不赋值
		if SmallVal > arr[i] {
			SmallVal = arr[i]
			SmallIndex = i
		}
	}
	fmt.Printf("最大值 %v 最大值下标 %v\n",MaxVla,Maxindex)
	fmt.Printf("最小值 %v 最小值下标 %v",SmallVal,SmallIndex)
}

func main()  {
	// 定义一个 arr 函数元素值为 11,22,33,44,55
	var arr [6]int = [...]int{11,2,33,34,20,89}
	
	// 调用 arr 函数并将 arr 数组传递过去
	Arr(&arr)
}