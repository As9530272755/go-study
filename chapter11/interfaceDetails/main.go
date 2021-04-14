package main

import "fmt"

type A interface {
	Say()
}

// 自定义 interger 类型，类型是 int
type interger int

func (i interger) Say() {
	fmt.Println("interger i = ",i)
}

func main()  {
	// 给 s 赋值为 interger 类型，然后把 10 传递给 i
	var s interger = 10

	// 定义 b 变量类型为 A interface，把 s 变量传递，因为 s 变量实现了 interger ，而 interger 类型中实现了 Say() 方法
	var b A = s

	// 调用 b 中的 Say()
	b.Say()
} 