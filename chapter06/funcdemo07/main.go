package main

import (
	"fmt"
)

// 累加器
func Adduppre() func (int) int {
	var n int = 10
	var str = "hello"
	return func (x int) int {
		n = n + x
		str = str + "a"
		fmt.Println("\nstr=",str)	// 第一次 helloa 第二次 helloaa 第三次 helloaaa
		return n
	}	
} 

func main() {

	// 将 Adduppre 函数赋值给 f 变量
	f := Adduppre()
	fmt.Println(f(1)) // 11 
	fmt.Println(f(2)) // 13
	fmt.Println(f(3)) // 16
}