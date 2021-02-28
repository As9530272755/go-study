package main

import (
	"fmt"
)

func main()  {
	var t int
	fmt.Println("输出需要打印得层数：")
	fmt.Scanln(&t)

	for i := 1 ; i <= t ; i++{
		for k :=1 ; k <= t - i ; k++{
			fmt.Printf(" ")
		}
		for j := 1 ; j <= 2 * i  - 1 ; j++{
			if j == 1 || j == 2 * i  - 1 || i == t {
			fmt.Printf("*")
			} else	{
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}