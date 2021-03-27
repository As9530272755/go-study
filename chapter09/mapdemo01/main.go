package main

import (
	"fmt"
)

func main()  {
	// map 的声明和注意事项

	// 这里声明了一个 a map
	var a map[string]string

	// 在使用 map 前需要先 make，make 的作用就是给 map 在内存中分配数据空间
	// 写法如下，make 小括号中先写 map 的数据类型，10 表示为给 a 分配了 10 个这样的 key、value 空间
	// 也就是说 a 可以放 10 对 key value 
	a = make(map[string]string,10)

	// 给这个 a 的 map 变量赋值 no1 的 key = songjiang
	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no3"] = "宋江"
	fmt.Println(a)
}