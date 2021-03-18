package main

import (
	"fmt"
)

func main()  {
	// string 这个数据类型，底层是一个 byte 数组，因此 string 也可以进行切片处理
	str := "hello&zgy"

	// 这时候想对 str 这个 string 类型进行一个切片处理，想把 &zgy 取出来
	// 此时就需要切片处理

	slice := str[5:]
	fmt.Println(slice)

	// 现在想把 str 字符串中的 h 改为 z
	
	// 把 str 转为为 byte 切片赋值给 arr
	arr := []byte(str)
 
	arr[0] = 'z'

	str = string(arr)

	fmt.Println(str)
	
	arr1 := []rune(str)

	arr1[0] = '张'

	str = string(arr1)
	fmt.Println(str)
}