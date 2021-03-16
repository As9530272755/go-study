package main

import (
	"fmt"
)

func main()  {
	// 演示切片 make 方式的使用
	// 对于切片而言，必须 make 才能够使用
	var slice []float64 = make([]float64,5,10)
	fmt.Printf("slice 默认值 %v\n",slice)

	// 给 slice 切片下标为 1 进行赋值为 10
	slice[1] = 10

	// 给 slice 切片下标为 3 进行赋值为 20
	slice[3] = 20

	fmt.Printf("赋值后的 slice 元素 = %v\n",slice)
	fmt.Printf("slice 的容量 = %v\n",cap(slice))
	fmt.Printf("slice 的 len = %v\n",len(slice))

	var slice1 []int = []int{1,3,5}
	fmt.Println("\n",slice1)
	fmt.Printf("slice1 的容量 = %v\n",cap(slice1))
	fmt.Printf("slice1 的 len = %v\n",len(slice1))

	var slice2 []int = make([]int,4)
	fmt.Printf("%v",&slice2[0])
}