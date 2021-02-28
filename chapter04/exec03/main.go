package main

import (
	"fmt"
)

func main ()  {
	
	// 求三个数的最大值
	
	// 先定义三个数
	var n1 int = 3
	var n2 int = 5
	var n3 int = 6
	var max int
	if n1 > n2 {
		max = n1
	} else if n1 > n3 {
		max = n1
	} else if n2 > n3 {
		max = n2
	} else {
		max = n3
	}
	fmt.Println("max 最大值为",max)
}