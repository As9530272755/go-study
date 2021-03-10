package main

import (
	"fmt"
)

func main()  {
	// 创建一个 byte 类型的 26 个元素的数组，分别放置 'A'-Z' 这 26 个英文字母 。
	// 使用 for 循环访问所有元素并打印出来。
	// 提示：字符数据运算 'A'+1-> 'B'
	
	// 思路：
	// 1.声明一个数组，长度为 26 类型为 byte
	// 2.使用 for 循环利用字符可以进行运算的特点来赋值 'A'+1-> 'B'

	// 代码:
	var myLetter [26]byte
	for i := 0 ; i < len(myLetter) ; i++ {
		// 'A' 是一个 byte 类型，所以这里要强制转 i 为 byte 类型
		myLetter[i] = 'A' + byte(i)

		// fmt.Printf 函数中的 %c 通过字符输出
		fmt.Printf("%c ",byte(myLetter[i]))
	}
}