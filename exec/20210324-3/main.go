package main

import (
	"fmt"
)

// 已知有个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
func main()  {
	// 定义一个升序的 arr 数组
	var arr [5]int = [...]int{1,22,33,44,55}

    // 在定义新的 arr1 数组，并且 arr1 长度比 arr 多一个
	var arr1 [6]int
    
    // 这里是给 arr1 index 最长的那个值赋值为 21
	arr1[len(arr)] = 21
	
    // for 循环遍历 arr 然后将 arr[i] 的值赋给 arr1[i]
	for i := 0 ; i < len(arr) ; i++ {
		arr1[i] = arr[i]
	}
	fmt.Println("赋值给 arr1 之前",arr1)
    
    // 冒泡算法为升序
    temp := 0
	for j := 0 ; j < len(arr1) -1; j++ {
		for i := 0 ; i < len(arr1) -1 ; i++ {
			if arr1[i] > arr1[i+1] {
				temp = arr1[i+1]
				arr1[i+1] = arr1[i]
				arr1[i] = temp
			}
		}
	}
	fmt.Println("排序后",arr1)
}