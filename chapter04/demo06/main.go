package main

import	(
	"fmt"
)

func main()  {
	// 演示 & 取地址符和 * 指针变量的使用范例

	// 1、演示 & 取地址符的范例
	// a := 10
	// fmt.Println("a 的地址=",&a)

	// // 2、演示 * 指针变量的使用范例
	// var ptr *int = &a
	// fmt.Println("ptr 指向的指针变量是=",*ptr)

	var n int
	var i int = 10
	var j int = 12

	if i > j { // 判断 i 是否大于 j ，如果大于 n 就等于 i
		n = i
	}	else {
		n = j // else 如果 i 不大于 j 就 n = j 因此 n = 12
    }
	fmt.Println("n=",n) // n = 12
}