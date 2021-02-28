package main

import (
	"fmt"
)

// 声明一个函数（该函数用来做测试）
func test()	bool{	// 使用 test 函数，并且这个函数返回值为真
	fmt.Println("test...")
	return true		// 定义 test 这个函数返回值为 true 真
}
func main()  {

	/*
	// 演示逻辑运算符的使用
	var age int = 40 // 给 age 赋值为 40
	if age > 30 && age < 50 { // 通过 if 做一个 与 判断，age 大于 30 并且小于 50 就输出 ok1
		fmt.Println("ok1")
	}

	if age > 30 && age < 40 {	// age 大于 30 并且 小于 40 就输出 OK2
		fmt.Println("OK2")
	}

	if age > 30 || age < 50 { // 通过 if 做一个 或 判断，age 大于 30 或者 age 小于 50 就输出 ok1 （满足条件）
		fmt.Println("ok3")
	}

	if age > 30 || age < 40 {	// age 大于 30 或者 小于 40 就输出 OK2（满足条件）
		fmt.Println("OK4")
	}


	if age > 30{	// 因为 age 赋值为 40 所以大于 30 条件成立为真输出 OK1
		fmt.Println("ok5")
	}

	if !(age > 30) {	// 小括号中 age=40 大于 30 条件为真，但是通过 ! 取反所以条件为假不打印 OK2
		fmt.Println("ok6")
	}
	*/

	// 演示短路与或者短路或运行的特点
	var i int = 10

	// if i > 9 && test(){ // i 大于 9 然后调用 test 函数。整个代码结果为 test 在输出 ok
	// 	fmt.Println("ok...")
	// }

	if i > 9 || test(){	// i=10 > 9 为真，第一给条件已经为真了所以不用再看 test 函数，并且将 test.. 、hello.. 进行输出
		fmt.Println("hello...")
	}
}