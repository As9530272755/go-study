package main

import (
	"fmt"
)

func main()  {
// 下面代码就是一个死循环。通常配合 break 语句使用。
	k := 1		// 给 k 赋值等于 1
	for {
		if k <= 10 {		// 然后判断 k 不小于等于 10 为真，大于 10 为假
		 fmt.Println("OK")	// 为真就执行该代码
		} else {			// 大于 10 就直接 else 语句
			break			// 然后 break 退出循环
		}
		k++					// 循环变量迭代
	 }
}