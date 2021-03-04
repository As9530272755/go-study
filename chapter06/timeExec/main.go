package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03()  {
	// 这里定义了 str 全局变量，并且赋值了空字符串
	str := ""

	// for 循环判断 hello + i 循环 10 万次
	for i := 0 ; i < 100000 ; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main()  {
	// 在执行 test03 前，先获取到当前的 unix 时间戳
	// 当然想要精准一点我们也可以让他拿到 unixnano 时间戳，这里我只是让他拿到 unix 时间戳就可以了
	
	// 这里给 start 变量赋值函数 time.Now().Unix() 也就是执行 main 函数的时间
	start := time.Now().Unix()
	
	// 调用了 test03 函数，这里就开始运行了 test03 函数里定义的内容
	test03()

	// 这里给 end 变量赋值了 time.Now().Unix() 也就是 main 函数执行完毕之后的 秒数
	end	:= time.Now().Unix()

	// 然后输出通过 end - start 变量的时间就能够得出跑完 test03 函数的总共用时
	fmt.Printf("执行 test03() 耗费时间为 %v 秒\n",end - start)
}