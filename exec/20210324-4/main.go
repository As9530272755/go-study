package main

import (
	"fmt"
)

func main()  {
	// 定义一个 3 行 4 列的二维数组，逐个从键盘输入值，编写程序将四周的数据清 0
	var arr[3][4]int

	// 输入值
	for i := 0 ; i < len(arr);i++ {
		for j := 0 ; j < len(arr[i]) ;j++ {
			fmt.Printf("请输入第 %d 行，第 %d 列的值\n",i+1,j+1)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Print("输入值后排序\n")
	// 输入值后排序
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0 ; j < len(arr[i]) ; j++ {
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}
	fmt.Print("清 0 后排序\n")
	// 清 0 后排序
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0 ; j < len(arr[i]) ; j++ {
			if i == 0 || i == len(arr) -1  || j == 0 || j == len(arr[i]) -1 {
				fmt.Print("0 ")
			} else {
				fmt.Print(arr[i][j]," ")
			}
		}
		fmt.Println()
	}
}