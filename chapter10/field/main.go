package main

import (
	"fmt"
)

// 如果结构体的字段类型是指针、slice 、map 的使用范例，如果没有赋值就全部为 nil 
// 如果需要使用这样的字段、需要先 make 才能够被使用

// 定义一个 Person 结构体
type Person struct{
	Name string
	Age int

	// 一共有 5 门课的成绩
	Scores [5]float64
	
	// 指针
	Ptr *int

	// 切片
	Slice []int

	// 在定义一个 map
	map1 map[string]string

}

// 这里又定义了一个 Monster 结构体
type Monster struct {
	Name string
	Age int
}

func main()  {

	// 定义一个结构体变量
	var p1 Person

	// 使用 slice 类型
	// 1.首先要 make
	// 2.对该 slice 赋值
	p1.Slice = make([]int,10)
	p1.Slice[0] = 100
	fmt.Println(p1.Slice)

	// 使用 map 类型
	// 1.首先要 make
	// 2.对该 map 进行赋值
	p1.map1 = make(map[string]string)
	p1.map1["no1"] = "zgy"
	p1.map1["no2"] = "haha"
	fmt.Println(p1.map1)

	// 定义两个 Monster
	var monster1 Monster
	monster1.Name = "白骨精"
	monster1.Age = 500

	// 这里再通过类型推导定义了一个 monster2 变量然后将 monster1 赋值给他
	monster2 := &monster1

	// 这里再通过类型推导定义了一个 monster3 变量然后将 monster1 赋值给他
	monster3 := &monster1
	// 因为结构体是一个值拷贝，所以这里 monster3.Name 改变之后并不会影响到 monster1 和 monster2 中的 name 字段
	monster3.Name = "孙悟空"

	fmt.Println("monster1 = ",monster1)
	fmt.Println("monster2 = ",monster2)
	fmt.Println("monster3 = ",monster3)
}