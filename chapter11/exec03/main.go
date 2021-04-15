package main

import (
	"fmt"
)

func main()  {
	// 先定义一个切片
	var intSlice = []int{12,-2,10,7,9,90}
	
	// 对 intSlice 进行排序
	// 使用冒泡算法对切片进行排序
	for i := 0 ; i < len(intSlice) -1; i++ {
		for j := 0 ; j < len(intSlice) -1; j++ {
			if intSlice[j] > intSlice[j+1] {
				intSlice[j],intSlice[j+1] = intSlice[j+1],intSlice[j]
			}
		}

	}
	fmt.Println("冒泡算法对切片进行排序：",intSlice)
}