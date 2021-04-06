package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

func main()  {
	var stu1 Student = Student{"tom",10}
	fmt.Println(stu1)

	stu2 := Student {"zzz",20}
	fmt.Println(stu2)

	var stu3 Student = Student {
		Name : "stu3",
		Age : 30,
	}
	fmt.Println(stu3)

	stu4 := Student {
		Name : "stu4!~",
		Age : 33,
	}
	fmt.Println(stu4)

	// 返回结构体的指针类型，在内存这个 stu5 先指向了一个地址，这个地址在指向了结构体数据空间
	var stu5 = &Student{"stu5~",100}
	stu6 := &Student{"stu6~",40}

	// 在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，不依赖字段的定义顺序
	var stu7 = &Student{
		Name : "stu7",
		Age : 77,
	}

	stu8 := &Student{
		Name : "stu8",
		Age :88,
	}

	fmt.Println(*stu5,*stu6,*stu7,*stu8)
}