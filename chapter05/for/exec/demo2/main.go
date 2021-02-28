package main

import (
	"fmt"
)

func main()  {
	var n int = 6	// 给 n 赋值为 6
	for i := 0 ; i <= n ; i++{	// i 的初始值为 0，并且小于等于 6 ，然后一次累计循环 6 次
		fmt.Printf("%v + %v = %v\n",i,n - i,n)	// i + n = n 也就是说加上循环次数最后等于的值为 6 的输出
	}
}