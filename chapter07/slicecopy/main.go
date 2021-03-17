package main

import (
	"fmt"
)

func main()  {

	// 定义切片 a ，int 类型，元素值 1，2，3，4，5
	var a []int = []int {1,2,3,4,5}
	
	// 通过 make 定义一个 slice1 的 int 类型切片，len 为 10
	var slice1 = make([]int,10)
	
	// 输出 slice1 由于没有赋值，默认为 10 个 0 ，因为他的 len = 10
	fmt.Println("拷贝前输出 slice1 = ",slice1)

	// 使用 copy 内置函数，将切片 a 的值拷贝给 slice1 切片
	copy(slice1,a)

	// 输出拷贝过后的 slice1 切片
	fmt.Println("拷贝后输出 slice1 = ",slice1)

	// 修改 a[0] = 999
	a[0] = 999

	// 输出 a 的值
	fmt.Println("修改后的 a = ",a)

	// 输出 slice1 的值
	fmt.Println("slice1 = ",slice1)
}