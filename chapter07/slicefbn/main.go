package main

import (
	"fmt"
)

func fbn(n int) ([]uint64) {
	// 声明一个切片，切片长度就为 n ，因为我们传入的是一个 n 肯定就有 n 个斐波那契
	// 通过 make 定义了 fbnSlice 切片，类型为 uint64 长度为 n
	fbnSlice := make([]uint64,n)
	
	// 第一个数和第二个数的斐波那契数都等于 1
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	// 其他的斐波那契数列通过 for 循环往这里面放
	for i := 2 ; i < n ; i++ {
		
		// 斐波那契数列第 n 个值 = 前两个斐波那契值相加
		fbnSlice[i] = fbnSlice[i -1] + fbnSlice[i -2]
	}

	// 返回值为 fbnSlice 切片
	return fbnSlice
}

func main()  {
	// 斐波那契数列切片练习题
	// 思路：
	// 1.声明一个函数 fbn(n int) ([]uint64),返回值为 []uint64 这种数据类型存放的值是最大的
	// 2.编写 fbn (n int) 进行 for 循环来存放斐波那契数列，0 和 1 的这两个对应的斐波那契值都为 1 需要单独处理
	
	// 测试,调用 fbn 函数并且赋值为 n = 10 ，也就是在 fbn 函数中的 fbnSlice 切片 len = 10
	// 并且在 main 函数中定义一个 MainfbnSlice 切片来接收 fbn 函数中的返回值 fbnslice 切片
	MainfbnSlice := fbn(10)
	fmt.Println(MainfbnSlice)
}