package main

import (
	"fmt"
)

// map 是引用类型，遵守引用类型传递的机制，如果在一个函数中接收 map ，修改后，会直接修改原来的这个 map

// 1.先定义一个函数,函数形参为 map1 是一个 map 类型
func modify(map1 map[int]int)  {
	
	// 这里将 map1 下标为 10 的值改为 900
	map1[10] = 900
}

func main()  {
	// 在 main 函数中声明一个 map
	map1 := make(map[int]int,2)
	
	// 给 map1 下标为 1，2，10，20 分别赋值
	map1[1] = 90
	map1[2] = 88
	map1[10] = 100
	map1[20] = 2

	// 调用 modify 函数，传入 map1
	modify(map1)
	
	fmt.Println(map1)
}