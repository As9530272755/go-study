package main

import (
	"fmt"
)

type MethodUtils struct {

}

func (m *MethodUtils) print(len int,width int) int {
	return len * width
}

func main()  {
	var mu MethodUtils
	area := mu.print(10,20)
	fmt.Println("该矩形面积为 = ",area)
}