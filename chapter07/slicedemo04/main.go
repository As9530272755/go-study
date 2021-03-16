package main

import (
	"fmt"
)

func main()  {
	var arr [5]int = [...]int{1,2,3,4,5}
	// 第一种写法
	var slice3 = arr[:3]
	fmt.Printf("第一种写法取出的 slice = %v\n",slice3)

	// 第二种写法
	var slice2 = arr[3:]
	fmt.Printf("第二种写法取出的 slice1 = %v\n",slice2)
	
	// 第三种写法
	var slice4 = arr[:]
	fmt.Printf("第三种写法取出的 slice2 = %v\n",slice4)

	// 这里 slice 已经是一个切片了
	var slice = arr[:]

	// 切片之后还可以继续切片
	slice1 := slice[1:2]
	
	// 对 slice1[0] 从新赋值为 100
	slice1[0] = 100
	fmt.Printf("slice1 = %v\n",slice1)
	fmt.Printf("slice = %v\n",slice)
	fmt.Printf("arr = %v\n",arr)
}