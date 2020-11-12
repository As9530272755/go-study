package main

import (
	"fmt"
	"strconv"
)

// 演示 golang 中 string 类型转换为基础数据类型
func main()  {
	var str string = "true"
	var b bool 

	// b , _= strconv.ParseBool(str) 说明：
	// 1、strconv.ParseBool(str) 该函数本身会返回两个值，第一个值为 bool 值，第二个值为 error
	// 2、这时候我关心的是 bool 值对这个 error 不关心所以我使用下划线来忽略第二个返回的 error 值
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v",b,b)

	var str2 string = "12345"
	var n1 int64 
	var n2 int

	// strconv.ParseInt(str2,10,0) 说明：
	// 1、 strconv.ParseInt(str2,10,0) 中的 str2 是取需要转换的 string 类型变量
	// 2、 10 表示为 10 进制
	// 3、 0 表示为 int 数据类型
	n1, _ = strconv.ParseInt(str2,10,0)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	fmt.Printf("n2 type %T n2=%v\n",n2,n2)

	var str3 string = "123.456"
	var f1 float64
	var f2 float32
	// f1, _ = strconv.ParseFloat(str3,64) 说明：
	// 1、strconv.ParseFloat(str3,64) str3 表示我们需要转的 string 类型变量为 str3，将他转换为 float 类型
	// 2、64 表示为我们要转换为 float64 数据类型
	// 3、f1, _  表示为我们也需要忽略 error 值
	f1, _ = strconv.ParseFloat(str3,64)
	f2 = float32(f1)	// 将 f1 转位 float32 数据类型，然后再赋值给 f2
	fmt.Printf("f2 type %T f2=%v\n",f2,f2)	// 输出 f2 类型和 f2 的值
	fmt.Printf("f1 type %T f1=%v",f1,f1)
}