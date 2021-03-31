package main

import (
	"fmt"
)

// 定义 Point 的结构体
type Point struct {
	x int
	y int
}

// 定义 Rect 结构体,定义了 leftUp,reghtDown 两个字段类型是 Point 结构体
type Rect struct {
	leftUp,reghtDown Point
}

func main()  {
	// 定义了 r1 变量，然后调用 Rect 结构体，leftUp 字段值为 Point 1,2 reghtDown 字段值为 Point 3,4	
	r1 := Rect{Point{1,2},Point{3,4}}

	// r1 有 4 个 int，在内存中是连续分布的
	// 打印地址
	fmt.Printf("r1.leftUp.x 地址 = %p \nr1.leftUp.y 地址 = %p \nr1.reghtDown.x 地址 = %p \nr1.reghtDown.y 地址 = %p\n",
	&r1.leftUp.x , &r1.leftUp.y , &r1.reghtDown.x , &r1.reghtDown.y)

}