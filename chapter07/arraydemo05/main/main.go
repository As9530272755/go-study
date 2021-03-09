package main

import (
	"fmt"
)

func main() {
	// 第一种方式，在定义数组的时候就初始化
	var numArr01 [3]int = [3]int{1,2,3}
	fmt.Println("numArr01 数组初始 = ",numArr01)

	// 第二种方式，通过类型推导，将赋值给该数组
	var numArr02 = [3]int{4,5,6}
	fmt.Println("numArr02 数组初始 = ",numArr02)

	// 第三种方式，直接不用写元素数量，通过 ... 的方式
	var numArr03 = [...]int{7,8,9}
	fmt.Println("numArr03 数组初始 = ",numArr03)

	// 第四种方式，通过下标加 : 加值的方式来定义.
	// 这里我定义了 0 的下标值为 800 ，1 的下标值为 200，2 的下标值为 1000
	var numArr04 = [...]int{0:800, 1:200, 2:1000}
	fmt.Println("numArr04 数组初始 = ",numArr04)

	// 第五种方式，通过类型推导的方式来初始化数组
	strArr05 := [...]string{"zhang","gui","yuan"}
	fmt.Println("strArr05 数组初始 = ",strArr05)
}