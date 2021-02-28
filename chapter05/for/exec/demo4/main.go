package main

import (
	"fmt"
)

func main()  {
	// 这里的 i 表示层数,一共有 9 层
	for i := 1 ; i <= 9 ; i++{
		for j := 1 ; j <= i ; j ++{ 	// 这里的 j 表示后面被乘以的变化的数
			fmt.Printf("%v * %v =%v\t",j,i,j *i) // 因为 j 是需要变化的所以就是 j * i 
		}
		fmt.Println()
	}
}