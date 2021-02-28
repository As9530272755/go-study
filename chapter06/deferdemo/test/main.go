package main

import (
	"fmt"
)

func sum(n1 int,n2 int) int {

	// 在将这句话压入到 defer 栈的时候是直接将 main 函数中 n1 = 10 的原相关值拷贝进 defer 栈
	defer fmt.Println("ok1 n1=",n1) 
	// 在将这句话压入到 defer 栈的时候是直接将 main 函数中 n2 = 20 的原相关值拷贝进 defer 栈
	defer fmt.Println("ok2 n2=",n2) 
 
	// 增加一句话
	n1++ // n1 ++ 后变为了 11
	n2++ // n2 ++ 后变为了 21

	res := n1 + n2
	fmt.Println("res=",res) // 输出 res = 32
	return res
}

func main()  {
	res2 := sum(10,20)			// 将 10 传递给 sum 函数的 n1，20 传递给 sum 函数的 n2
	fmt.Println("res2 = ",res2) // 最后输出 32 因为 return 的值是 32 ，所以将返回值传递给 res2 变量
}