package main

import (
	"fmt"
)

// 这里 n1 就是 *int 类型
func test(n1 *int)  {
	*n1 = *n1 + 10
	fmt.Println("test(n1)=",*n1)
}

func main()  {
	num := 20
	test(&num)
	fmt.Println("main(num)=",num)
}