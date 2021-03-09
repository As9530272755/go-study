package main

import (
	"fmt"
)

func main()  { 
	// 定义一个数组,下标为 5 保存 5 个成绩，该数组类型为 float64
	var score [5]float64

	// 写一个 for 循环，在终端循环输入值
	for i := 0 ; i < len(score) ; i++{
		// 通过 printf 在终端输入
		fmt.Printf("请输入第 %d 个的元素值\n",i+1)
		
		// 通过 scanln 接收 score[i] 的值，并存入给 score 数组的这 5 个元素中
		fmt.Scanln(&score[i])
	}

	// 遍历数组，将数组进行打印
	for i := 0 ; i < len(score) ; i++ {
		// printf 的 %d 用来接收 i 的值，然后再将刚才 scanln 存入到 score 数组中的元素 i 进行打印
		fmt.Printf("score[%d] = %v\n",i,score[i])
	}
}