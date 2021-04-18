package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main()  {
	var a interface{}
	var point Point = Point{1,2}
	
	a = point // 能赋值，a 是一个空接口类型，能够接收任何变量

	
	var b Point

	// b 直接赋值给 a 是不可以
	// b = a.(Point) 就可以这个就叫做 类型断言，a.(Point) 就是将这个 a 转换为 Point 结构体类型赋给 b
	b = a.(Point)
	fmt.Println(b)

	// 类型断言的其他案例

	// 定义 x 为 空接口类型
	var x interface {}
	var c float32 = 1.1

	// 空接口可以接收任意类型，把 c 赋值给 x，此时 x 为 float32 类型
	x = c

	// x 转为 float32 [使用类型断言]
	y,ok := x.(float32)
	if ok {
		// 输出转换成功
		fmt.Println("convert success")
	} else {
		fmt.Println("convert fail")
	}
	fmt.Printf("y 的类型 %T 值= %v",y,y)
}