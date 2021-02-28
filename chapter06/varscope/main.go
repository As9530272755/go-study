package main

import (
	"fmt"
)

var age int = 50 // age 由于首字母是小写的只能够在整个包里面使用
var Name string = "jack" // Name 由于首字母是大写不但在该包能够使用在其他包中也能使用

func test()  {
	// age 和 Name 变量的作用域就只在 test 函数内部
	age := 10
	Name := "tom"
	fmt.Println("test()age=",age)	// 10
	fmt.Println("test()Name=",Name)	// tom
}

func main()  {
	fmt.Println("age=",age)	// 输出 50 调用全局变脸中的 age
	fmt.Println("Name=",Name) // 输出 jack 调用全局变脸中的 Name
	test()	// 调用 test 函数


	var i int 
	// 在 for 循环中定义了 i 这个变量，这个 i 只能够在当前 for 代码块中使用
	for i = 0 ; i <= 10 ; i++ {
		fmt.Println("i=",i)
	}
	
	fmt.Println("i=",i)


}