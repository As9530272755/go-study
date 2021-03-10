package main

import (
	"fmt"
)

func main()  {
	// 请求出一个数组的最大值，并得到对应的下标。
	
	// 思路：
	// 1.声明一个数组 var intArr[5] = [...]int{1,-1,9,90,11}
	// 2.假定第一个元素就是最大值，下标就是 0 
	// 3.然后从而第二元素开始循环比较，如果有更大的数则交换

	var intArr [6]int = [...]int{20,-1,9,90,11,100}

	// 定义变量，假定为 intArr[0] 是最大值
	maxVal := intArr[0]
	
	// 定义变量，假定 0 为最大值的下标
	maxValIndex := 0

	for i := 1 ; i < len(intArr) ; i++ {
		// if 判断然后从而第二元素开始循环比较，如果有更大的数则交换
		if maxVal < intArr[i] { // intArr[i] 就是当前不停循环比较的值
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal = %v\nmaxValIndex = %v",maxVal,maxValIndex)
}