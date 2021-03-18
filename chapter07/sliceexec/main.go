package main

import (
	"fmt"
)

func test(slice []int)  {
	slice[0] = 100
}

func main()  {
	// 创建一个 a 的切片
	var a []int = []int{1,2,3,4,5}
	
	// 通过 make 定义了 slice 切片 len 为 1
	var slice4 = make([]int,1)
	fmt.Printf("最初的 slice = %v\n",slice4)
	
	// 这里将 a 切片的值赋值给 slice
	copy(slice4,a)
	fmt.Printf("拷贝后的 slice = %v\n",slice4)

	var slice1 []int
	var arr [5]int = [...]int{1,2,3,4,5}
	slice1 = arr[:]
	var slice2 = slice1
	slice2[0] = 10

	fmt.Println("slice2\n",slice2) // 输出10，2，3，4，5
	fmt.Println("slice1\n",slice1) // 输出10，2，3，4，5
	fmt.Println("arr\n",arr)		 // 输出10，2，3，4，5
	fmt.Println()

	var slice3 = []int {1,2,3,4}
	fmt.Println("slice3 = ",slice3) // 输出 1，2，3，4

	test(slice3)
	fmt.Println("slice3 = ",slice3) // 输出 100，2，3，4
}