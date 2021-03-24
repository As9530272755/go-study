package main

import (
	"fmt"
)

// 已知有个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
func main()  {
	var arr [5]int = [...]int{1, 3, 4, 5, 6}
	var arr1 = [6]int{}
	var aa = 2
	ai := 0
	for i := 0; i < len(arr); i++ {
	 arr1[ai] = arr[i]
	 if arr[i] < aa && arr[i+1] > aa {
	  arr1[i+1] = aa
	  ai = ai + 1
	 }
	 ai++
	 fmt.Println(i)
	}
	fmt.Println(arr1)
}