package main

import (
	"fmt"
)

func main()  {
	// 演示二维数组的遍历
	arr3 := [2][3]int{{1,2,3,},{4,5,6}}

	// for 循环进行遍历
	// 先遍历外层一维数组，len(arr3) = 2
	for i := 0 ; i < len(arr3) ; i++ {

		// 在遍历里面这个一维数组，len(arr3[i]) = arr[i] = 3 个元素
		for j := 0 ; j < len(arr3[i]); j++ {

			// 输出一个完整的二维数组 arr3[i][j]
			fmt.Printf("%v\t",arr3[i][j])
		}

		// 每遍历一次就进行就进行换行
		fmt.Println()
	}
	fmt.Println()
	// for-range 的方式进行遍历
	// i = arr[2] 一维数组
	// v = arr[2][3] 一维数组 {1,2,3,4,5,6}
	for i , v := range arr3 {
		for j,v1 := range v {
			fmt.Printf("arr3[%v][%v] = %v\t",i,j,v1)
		}
		fmt.Println()
	}
}