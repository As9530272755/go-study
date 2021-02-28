package main

import (
	"fmt"
)

func main()  {
	var char byte
	fmt.Println("请输入对应的 a b c d e")
	fmt.Scanf("%c",&char)
	switch char{
		case 'a' :
			fmt.Println("A")
		case 'b' :
			fmt.Println("B")
		case 'c' :
			fmt.Println("C")
		case 'd' :
			fmt.Println("D")
		case 'e' :
			fmt.Println("E")
		default :
			fmt.Println("outher")
	}
}