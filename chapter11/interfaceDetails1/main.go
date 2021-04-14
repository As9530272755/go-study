package main

import "fmt"

type A interface {
	Say()
}

type Stu struct {
	Name string
}

func (s Stu) Say() {
	fmt.Println("Stu.Say()")
}

func main()  {
	var sa Stu // Stu 结构体变量 sa ，实现了 Say() 方法，也就实现了 A interface 中的 Say()
	
	// 然后再将 A interface 类型传递给 a ，并且将 sa 结构体变量传递给 a 变量
	var a A = sa
	a.Say()
} 