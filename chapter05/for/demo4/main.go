package main

import (
	"fmt"
)

func main()  {
	// 字符串通过 for-range 的方式进行遍历
	var str string = "abc~ok上海"
	for index, val := range str {
		// 在取值的时候 index 就是通过按照顺序进行取值，由 0 开始依此类推
		fmt.Printf("index=%d, val=%c\n",index , val)
	}
} 