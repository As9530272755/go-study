package main

import (
	"fmt"
)

/*
通过函数递归的方式求出斐波那契
给你一个整数 n ， 求出他的斐波那契数是多少？
*/

func fbn(n int ) int {
	fmt.Println(n)
	if n == 1 || n == 2 {
		return 1 
	} else {
		a := fbn( n - 1 )
		b := fbn( n -2 )
		return a+b 
	}
	 
}

// 5
// 4
// 3
// 2 = 1

// 3 = 1
// 4 = 1
// 5 = 1 1
// 6 = 1 1
// 4 = 1 


func main()  {
	// 测试
	res := fbn(6)
	fmt.Println("res=",res)
}