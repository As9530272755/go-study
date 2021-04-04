package main

import (
	"fmt"
)

/*
1. 声明一个结构体 Circle（圆） ，字段为 redius（半径）
2. 声明一个方法 area（面积） 和 Circle 绑定，可以返回面积
*/
type Circle struct {
	redius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.redius * c.redius
}

func main()  {

	// 创建一个 circle 结构体类型变量
	var c Circle
	c.redius = 4.0
	res := c.area()
	fmt.Println("面积是 = ",res)
}