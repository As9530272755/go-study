package main

import (
	"fmt"
)

func chenfa(n int)  {
	// 表示层数
	for i := 1 ; i <= n ; i++ { 
		// 表示每次每层需要乘的数
		for j := 1 ; j <= i ; j++ {
			fmt.Printf("%v * %v = %v \t",j,i,j * i)
		}
		fmt.Println()
	}
}

func main()  {
	var n int
	fmt.Println("请输入需要打印的乘法表层数：\t")
	fmt.Scanln(&n)
	chenfa(n)
}