package main

import (
	"fmt"
)

// 全局匿名函数

// 通过 var 定义一个全局变量 Func1 ，并且将 func 这个匿名函数赋值给 Func1 此时此刻这个 Func1 就是一个全局匿名函数
var (
	// Func1 就是一个全局匿名函数，并且首字母是需要大写的，行参事 n1 和 n2 并且都是 int 类型，返回列表也是 int 类型
	Func1 = func (n1 int,n2 int) int {
		// 返回值为 n1 * n2
		return n1 * n2
	}
)

func main()  {
	// 然后在 main 函数中使用 Func1 这个全局匿名函数，并且给行参赋值为 n1=2 n2=5
	res := Func1(2,5)
	fmt.Println("res=",res)
}