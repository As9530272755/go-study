package main

import (
	"fmt"
	"math"
)

func main()  {
	// 分析思路
	// 1.a b c 是三个 float64
	// 2.使用到给出的数学公式 提示1：x=(-b+sqrt(b2-4ac))/2a 提示2：×2=(-b-sqrt(b2-4ac))/2a
	// 3.使用到多分支
	// 4.使用到 math.Squr 方法 使用手册

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0
	m := b * b - 4 * a * c

	// 多分支判断
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v",x1,x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v",x1)
	} else {
		fmt.Println("无解")
	}
}