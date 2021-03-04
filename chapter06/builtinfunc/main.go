package main

import (
	"fmt"
)

func main()  {
	
	// 第一种方式：通过 fmt.Printf 的方式获取一个值得类型和它在内存中的地址
	num1 := 100
	fmt.Printf("num1 的类型 = %T\nnum1 的值 = %v\nnum1 的内存地址 = %v\n",num1,num1,&num1)

	// 第二种方式：
	num2 := new(int)	// 通过这样的分配其实是一个指针类型

	// num2 的类型 	 = *int
	// num2 的值   	 = 指向内存空间的地址
	// num2 内存地址 = num2 内存本身地址
	
	// 通过 num2 找到了指向的值 0 ，并且我们将这个默认值改为了 2
	*num2 = 2
	fmt.Printf("num2 的类型 = %T\nnum2 的值 = %v\nnum2 的内存地址 = %v\nnum2 指向的值 = %v",
	num2,num2,&num2,*num2)
} 