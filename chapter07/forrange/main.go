package main

import (
	"fmt"
)

func main()  {
	// 演示 for-range 遍历数组

	var heroes [3]string = [3]string{"宋江","吴用","卢俊义"}

	// value 就是每个数组对应的 元素，然后 _ 是一个占位符，也就是说忽略 index
	for _,value := range heroes{
		// 打印 value
		fmt.Printf("value = %v\n",value)
	}
}