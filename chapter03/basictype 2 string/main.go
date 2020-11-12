package main

import (
	"fmt"
	"strconv" // 引入 strconv 包
)

// 演示 golang 中基本数据类型转成 string 使用
func main()  {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的 str

	// 使用第一种方式来转换 fmt.Sprintf

	str = fmt.Sprintf("%d",num1) // 意思是我将 num1 转成一个 string 类型然后返回给 str 这个变量
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type %T str=%v\n",str,str)

	str = fmt.Sprintf("%t",b)
	fmt.Printf("str type %T str=%q\n",str,str)
	str = fmt.Sprintf("%c",myChar)  // %c 是按照字符来进行输出的，然后 mychar 给我转换成一个字符串
	fmt.Printf("str type %T str=%v",str,str)

	// 使用第二种方式来进行转换 strconv 包中的函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3),10) // 10 表示转成 10 进制,并且将 num 3 由 int 类型转为 int64
	fmt.Printf("str type %T str=%q\n ",str,str)


	// strconv.FormatFloat(num4,'f',10,64)
	// 说明：'f' 表示转成的格式。10 表示小数位保留 10 位也就是小数点之后可以有 10 位数。64 表示 num4 这个变量位 float64 位的
	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type %T str=%q\n",str,str)

	str = strconv.FormatBool(b2)	// 通过 strconve.FormatBool 函数直接调用 b2 的赋值
	fmt.Printf("str type %T str=%q\n",str,str)
} 