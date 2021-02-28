package main

import (
	"fmt"
) 

func main()  {
	// 输出 10 你好、尚硅谷
	// i 等于循环变量并且定义为 1 ，然后小于或者等于循环 10 次，并且是 i++ 一次累计
	for i := 1; i <= 10; i++ {
		fmt.Println("你好、尚硅谷",i)
	}
}