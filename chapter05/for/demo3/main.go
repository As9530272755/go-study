package main

import (
	"fmt"
)

func main()  {
	var str string = "hello world! 北京"
	str2 := []rune(str)	// 我们将上面 str 变量的 string 类型强制转为切片，这的小括号调用的是上面的 str 变量
	for i := 0 ; i < len(str2) ; i++{
		fmt.Printf("%c\n",str2[i]) // 使用到了下标。。
	}
}