package main

import (
	"fmt"
)

// 给 person 结构体添加 speak 方法，输出 xxx 是一个好人

// 定义一个 person 结构体
type Person struct {
	Name string
}

func (p Person) getSum(n1 int,n2 int) int {
	sum := n1 + n2
	return sum
}

func main()  {
	var p Person
	p.Name = "掌柜元"
	sum := p.getSum(10,20)

	fmt.Println(p.Name,"两个数相加放回的结果为 ",sum)
}