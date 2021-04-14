package main

import (
	"fmt"
)

// 定义一个 interface 名为 T ，不给它赋任何方法
type T interface {
}

type Stu struct {

}

func (s Stu) Say() {
	fmt.Println("Say..")
}

func main()  {
	var stu Stu
	var t T = stu
	fmt.Println(t)
	
	// 还有一种写法如下，直接给 t2 这个变量赋值为 interface 类型，传递 stu 结构体给 t2
	var t2 interface{} = stu
	fmt.Println(t2)

	// 定义 num1 的变量赋值 8.0
	num1 := 8.0

	// 将 num1 变量赋值给 t2
	t2 = num1
	fmt.Println(t2)
}