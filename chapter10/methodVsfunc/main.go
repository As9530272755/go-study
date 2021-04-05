package main

import (
	"fmt"
)

// 定义一个 Person 结构体
type Person struct {
	Name string
}

// 定义 test01 函数，形参为 p 类型为 person 结构体
// 结论：如果是函数，Person 为值类型就不能够将指针传递，Person 为指针类型就只能传地址
func test01(p Person)  {
	fmt.Println(p.Name)
}

// 反正，这里我定义的是 test02() 函数，传入的是 *Person 
func test02(p *Person){
	fmt.Println(p.Name)
}

func (p *Person) test03() {
	p.Name = "jack"
	fmt.Println("test03",p.Name)
}

func main() {
	// 定义 p 变量类型为 Person 同时赋值 tom
	p := Person{"tom"}
	
	// 调用 test01 函数，将 p 变量的值传递过去
	test01(p)

	// 调用 test02 函数，将 p 的地址传递过去
	test02(&p)
	p.test03()
	fmt.Println("main",p.Name)
}