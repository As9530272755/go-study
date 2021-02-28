package main

import (
	"fmt"
)

func main()  {

	a := 10
	b := 20

	b = a + b // a = a(10) + b(20) = 30 ，也就是 a + b 的总和
	a = b - a // b = a(30) - b(20) = 10 ，因为此时 a = 30 继承上一个变量赋值
	b = b - a // a = a(30) - b(10) = 20 ，b 继承了上一条命令的计算此时为 10。
	fmt.Printf("a = %v\nb = %v",a,b)
}