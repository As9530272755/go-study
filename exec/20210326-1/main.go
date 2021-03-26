package main

import (
	"fmt"
)

func main()  {
	// 定义一个4行4列的二维数组，逐个从键盘输入值
	// 然后将第1行和第4行的数据进行交换，将第2行和第3行的数据进行交换
	
	// 定义 4 行 4 列得二维数组
	var arr[4][4]int
	
	// 遍历该 二维数组并进行赋值
	for i := 0 ; i < len(arr) ;i++ {
		for j := 0 ;j < len(arr[i]) ; j++ {
			fmt.Printf("输入第 %d 行 %d 列\n",i+1,j+1)
			fmt.Scanln(&arr[i][j])
		}
		fmt.Println()
	}

	// 定义两个变量，temp 表示第一行，temo1 表示第二行
	temp := arr[0]
	temp1 := arr[1]
	for i := 0 ; i < len(arr) ; i++ {
		for j := 0 ; j < len(arr[i]) ; j++ {
		}

		// 如果当 i = 0 也就是第一行得时候，第1行和第4行的数据进行交换
		if i == 0 {
			temp = arr[0]
			arr[0] = arr[3]
			arr[3] = temp
		}
				
		// 如果当 i = 1 也就是第二行得时候，第2行和第3行的数据进行交换
		if i == 1 {
			temp1 = arr[1]
			arr[1] = arr[2]
			arr[2] = temp1
		}

		// 输出arr[i],把第一个一维数组进行输出
		fmt.Print(arr[i]," ")
		fmt.Println()
	}
}