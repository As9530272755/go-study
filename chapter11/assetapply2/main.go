package main

import (
	"fmt"
)

// 定义一个 Student 结构体
type Student struct {
	name string
}

// 编写一个函数用来判断输入的参数是什么类型
// 定义形参为 items... 可以接收多个变量，然后类型为 interface{} 空接口类型，可以接收任意类型
func TypeJudge(items... interface{}) {
	for i,v := range items {
		switch v.(type) {
			case bool:
				fmt.Printf("第 %v 参数是 bool 类型，值是 %v\n",i,v)
			case float32:
				fmt.Printf("第 %v 参数是 float32 类型，值是 %v\n",i,v)
			case float64:
				fmt.Printf("第 %v 参数是 float64 类型，值是 %v\n",i,v)
			case int,int32,int64:
				fmt.Printf("第 %v 参数是 整数 类型，值是 %v\n",i,v)
			case string:
				fmt.Printf("第 %v 参数是 string 类型，值是 %v\n",i,v)
			case Student:
				fmt.Printf("第 %v 参数是 Student 类型，值是 %v\n",i,v)
			case *Student:
				fmt.Printf("第 %v 参数是 *Student 类型，值是 %v\n",i,v)
			default :
				fmt.Printf("第 %v 参数类型不确定，值是 %v\n",i,v)
		}
	}
}

func main()  {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300

	// stu1 变量为 Student 类型
	stu1 := Student{"tom"}

	// stu2 变量为 &Student 类型
	stu2 := &Student{}

	// 调用函数判断这几个变量到底是什么类型
	TypeJudge(n1,n2,n3,name,address,n4,stu1,stu2)
}