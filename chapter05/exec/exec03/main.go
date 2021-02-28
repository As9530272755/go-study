package main

import (
	"fmt"
)

func main()  {
	var n1 int32 = 10
	var n2 int32 = 5
	if (n1 + n2) % 3 == 0 && (n1 + n2) % 5 == 0 {
		fmt.Println("能够被 3 和 5 整除")
	} else {
		fmt.Println("不能被整除")
	}
}