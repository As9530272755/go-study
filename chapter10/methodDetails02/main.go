package main

import (
	"fmt"
)

/*
Golang 中的方法作用在指定得数据类型上（即：和指定得数据类型绑定），
因此自定义类型都可以有方法，而不仅仅是 struct(结构体)，比如 int、float32 等都可以有自己的方法
*/

// 给 int 做了一个自定义，相当于 integer 就是 int 的别名
type integer int

// 给他绑定方法,这个方法叫做 print()
func (i integer) print() {
	fmt.Println("i =",i)
}

// 请编写一个方法，可以改变 i 的值
func (i *integer) change() {
	*i = *i +1
}

func main()  {
	// 定义一个 maini 变量，类型为 integer 类型 值为 10
	var maini integer = 10

	// 通过 maini 变量调用 print 方法，将 maini 值输出
	maini.print()

	maini.change()
	fmt.Println("修改后的 change 方法中的 i =",maini)
}