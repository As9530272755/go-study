package main

import (
	"fmt"
)

type MethodUtils struct {
	// 没有个这个结构体添加字段也是可以的
}

// 给 MethodUtils 结构体编写方法
func (ma *MethodUtils) print(m int,n int) {
	for i := 0 ; i < m ; i++ {
		for j := 0 ; j < n ;j++ {
			// 由于不换行打印就通过 print
			fmt.Print("*")
		}
		// 内循环输出完一次就换行，将其输出 10 次内循环
		fmt.Println()
	}
}

func main()  {
	// 在 main 函数中调用该方法
	// 定义
	var mu MethodUtils
	mu.print(4,3)
}