package main

import (
	"fmt"
)

func main()  {
	// 通过该代码来看一下数组的内存布局
	
	// 当我们定义完数组后其实，其实数组的各个元素已经有默认值了，默认值就是 0
	var intArr [3]int
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr 的地址 = %p\n",&intArr)
	fmt.Printf("intArr[0] 第一个元素地址 = %p\n",&intArr[0])
	fmt.Printf("intArr[1] 第二个元素地址 = %p\n",&intArr[1])
}