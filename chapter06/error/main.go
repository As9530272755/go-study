package main

import (
	"fmt"
)

func test()  {
	// 使用 defer + recover 来捕获和处理异常,这里编写了一个匿名函数
	// defer ： 当执行 test 函数完毕之后就会调用 defer 中的 func 匿名函数
	//然后将异常传入给 recover 函数在赋值给 err 我们在进行 if 判断
	defer func ()  {
		err := recover() 	// recover 是一个内置函数可以捕获到异常，这里赋值给了 err
		if err != nil {		// nil 是一个 0 值，如果 err 这个变量不等于 0 时就为错误
			fmt.Println("err = ",err)
			// 这里就可以将错误信息发送给管理员,如下的伪代码：
			fmt.Println("发送邮件至 admin@qq.com")
		}
	}()	// 匿名函数结尾，因为没有值所以这里用一个 () 即可，也就是将这个匿名函数压入到 defer 栈中
	num1 := 10
	num2 := 0
	res := num1 / num2	// 异常会在 test 函数这个位置发生
	fmt.Println("res = ",res)
}


func main()  {
	// 测试，调用 test 函数
	test()
	// 当 test 函数中的 defer 处理完了异常之后下面的 fmt 函数任然会执行
	fmt.Println("main()下面的代码...")
}