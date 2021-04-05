package main

import (
	"fmt"
)

type MethodUtils  struct {
	Num int
}


func (ma MethodUtils) Get() {
	for i := 1;i <= ma.Num ; i++ {
		for j := 1 ; j <= i ; j++ {
			fmt.Printf("%v*%v=%v\t",j,i,i * j)
		}
		fmt.Println()
	}
}

func main()  {
	var mu MethodUtils
	fmt.Println("请输入打印乘法表的列值")
	fmt.Scanln(&mu.Num)
	mu.Get()
}