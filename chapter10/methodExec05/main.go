package main

import (
	"fmt"
)

type MethodUtils struct {

}

// 根据行、列、字符打印对应行数和列数的字符，比如:行:3，列:2，字符*，则打印相应的效果
func (ma *MethodUtils) print(n int,m int,key string) {
	for i := 0 ; i < n ; i++ {
		for i := 0;i < m ;i++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main()  {
	var mu MethodUtils
	mu.print(3,2,"*")
}