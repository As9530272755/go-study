package main

import (
	"fmt"
)

// 演示 golang 中指针类型
func main()  {
	
	// 基本数据类型在内存中的布局
	var i int = 10
	
	// 我们怎么将 i 的地址取出
	fmt.Println("i 的地址 = ",&i)

	// var ptr *int = &i 表示的含义有以下几点
	// 1、 ptr 是一个指针变量
	// 2、 ptr 的类型是一个 *int ，也就是说它是指向一个 int 类型的一个指针
	// 3、 ptr 本身的值为 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr 的地址 =%v", &ptr)

	// 使用 *ptr 也就是说告诉编译器，将 ptr 指针指向的那个值取出，就是 &i 的值
	// 实际上第一步还是将 ptr 存放的地址取出来了，然后这里的 * 就把这个地址所对应的值返回输出
	fmt.Printf("ptr 指向的值 = %v",*ptr) 
}