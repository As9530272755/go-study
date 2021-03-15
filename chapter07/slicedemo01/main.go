package main

import (
	"fmt"
)

func main()  {
	
	// 演示切片的基本使用
	// 先定义一个数组
	var intArr[5]int = [...]int{1,22,33,66,99}
	fmt.Printf("修改前的 intArr 第 2 个下标 = %v\n",intArr[2])
	// 定义一个切片
	// 下面切片的说明
	// 1.slice 就是切片名
	// 2.intArr[1:3] 表示 slice 引用 intArr 数组 
	// 3.引用intArr数组的起始下标为 1 但是不包含 3 这个下标。也就是说只有 1 2 这两个下标
	// 4.也就是说 slice = 22，33
	slice := intArr[1:3]
	fmt.Printf("intArr 数组本身 = %v\n",intArr)
	fmt.Printf("slice 切片的元素 = %v\n",slice)
	fmt.Printf("slice 切片的元素个数 = %v\n",len(slice))

	// 容量：就是当前这个 slice 切片目前可以存放的最大个数的元素
	// 这里调用了 cap 内置函数来求出 slice 切片的容量
	fmt.Printf("slice 切片的容量 = %v\n",cap(slice))

	slice[1] = 34
	fmt.Printf("修改后的 intArr 第 2 个下标 = %v",intArr[2])
}