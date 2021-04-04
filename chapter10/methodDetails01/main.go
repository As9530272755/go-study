package main

import (
	"fmt"
)

type Circle struct {
	redius float64
}

// 为了提高效率，通常我们方法和结构体的指针类型绑定，也就是通过 * 取值符号，这时就是传入 c 的地址
func (c *Circle) area() float64  {
	
	// 这里将 c.redius 改为了 10
	c.redius = 10
	fmt.Printf("area() c 的地址 = %p\n",c)
	return 3.14 * c.redius * c.redius
}

func main()  {
	var c Circle
	c.redius = 4.0
	res := c.area()
	
	// 因为是传入的地址，所以在 area() 方法中改变了 c.redius 的值也会直接改变 main 函数中 c.redius的值
	fmt.Println("圆的面积 = ",res)
	fmt.Printf("main() c 的地址 = %p",&c)
}