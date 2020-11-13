package main

import (
	"fmt"
)

func main()  {
	var num int = 9
	fmt.Println(" num 地址是 ",&num)

	// 将 ptr 指针指向 num 的内存空间，也就是说现在 ptr 就是 num
	var ptr *int = &num
	
	// 然后修改 ptr 为 10 ，间接就将 num 改为了 10
	*ptr = 10

	// 通过 fmt 包中的 Println 函数输出 num 的值
	fmt.Println("num = ",num)
}