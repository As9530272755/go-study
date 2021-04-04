package main

import (
	"fmt"
)

type MethodUtils struct {
	// 没有个这个结构体添加字段也是可以的
}

// 给 MethodUtils 结构体编写方法
func (m *MethodUtils) print() {
	for i := 0 ; i < 10 ; i++ {
		for j := 0 ; j < 8 ;j++ {
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
	mu.print()
}