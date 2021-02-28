package main

import (
	"fmt"
)

func sum(n1 int, args...int) int {
	sum := n1
	// 遍历 args
	for i := 0 ; i < len(args); i++ {
		sum += args[i]	// args[0] 表示取出 args 切片的第一个元素的值，其他以此类推
	}
	return sum
}

func main()  {
	// 测试一下可变参数的使用
	res1 := sum(10,0,-1,90,10)
	fmt.Println("res1 = ",res1)
}