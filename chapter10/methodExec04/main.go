package main

import (
	"fmt"
)

type MethodUtils struct {

}

func (m *MethodUtils) JudgeNum(n int) {
	if n %  2 == 0 {
		fmt.Println(n,"是偶数")
	} else{
		fmt.Println(n,"奇数")
	}
}

func main()  {
	var mu MethodUtils
	mu.JudgeNum(10)
}