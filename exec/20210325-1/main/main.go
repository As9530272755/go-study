package main

import (
	"fmt"
)

func main()  {
	var arr[3][4]int
	for i := 0 ;i < len(arr) ; i++ {
		for j := 0;j < len(arr[i]) ; j++ {
			fmt.Printf("请输入第 %d 行，第 %d 列的数\n",i+1,j+1)
			fmt.Scanln(&arr[i][j])
		}
		fmt.Println()
	}
	
	for i := 0 ; i < len(arr) ;i++ {
		for j := 0 ; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr) -1 || j == 0 || j == len(arr[i]) -1 {
				fmt.Print("0 ")
			} else {
				fmt.Print(arr[i][j]," ")
			}
		}
		fmt.Println()
	}
}