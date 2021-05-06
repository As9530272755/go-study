package main

import (
	"fmt"
)

// 编写一个被测试的函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

func Sub(n int, n1 int) int {
	return n - n1
}

func main() {
	// 传统的测试方法，在 main 函数中使用看看结果是否正确
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper 错误，返回值 = %v 期望值 = %v", res, 55)
	} else {
		fmt.Println("测试正确", res)
	}

}
