package main

import (
	"fmt"
)

func test(char byte) byte {
	return char + 1
}

func main()  {

	var n1 int32 = 51
	var n2 int32 = 20
	switch n1 {	// 也就是说这个 n1 = n2 或者等于 10 或者 5 都能够匹配到就能输出 ok1
		case n2 , 10 , 5 :
			fmt.Println("ok1")
		case 92 :
			fmt.Println("ok2")
	}
}