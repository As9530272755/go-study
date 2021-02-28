package main

import (
	"fmt"
)

func test()  int {
	return 90
}

func main()  {
	// 常规赋值运算符使用演示
	// var i int
	// i = 10 // 最基本赋值案例
	// fmt.Println("i=",i)

	// 有两个变量，a 和 b ，要将其进行交换，最终打印结果

	// 原值 a = 9 , b = 2 。经过处理以后得到 a = 2 ,b = 9
	a := 9
	// b := 2
	// fmt.Printf("交换前的情况是 a= %v,b=%v\n",a,b)
	
	// 定义一个临时变量 t
	// t := a // 将 a 的值保存到 t 临时变量中
	// a = b  // 再将 b 赋值给 a，这时 a = b
	// b = t  // 然后在把 t 赋值给 b，从而 b = a
	// fmt.Printf("交换后的情况是 a= %v,b=%v",a,b)

	// // 复合赋值的操作
	// a += 7 // 等价于 a = a + 7 = 9 + 7 = 16
	// fmt.Println("a=",a)

	// var c int
	// c = a + 3 // 将等号右边的 a + 3 先进行计算然后再赋值给 c ， c = a + 3 = 9 + 3 = 12
	// fmt.Println("c=",c)

	// 举例说明
	var b int 
	var c int
	var d int
	var e int
	b = a // 等号的右边可以是一个变量
	c = 8 + 2 * 7 // 等号的右边可以是一个表达式
	d = test() + 90 // 等号右边是一个函数,因为这个 test() 函数返回值为 90 所以是一个表达式，并且还能将 test() 函数的返回值再进行二次计算都没问题
	e = 890 // 等号的右边可以是一个常量 890 就是一个 常量
	
	fmt.Printf("b=%v\nc=%v\nd=%v\ne=%v",b,c,d,e)


}