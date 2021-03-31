package main

import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func main()  {
	// 方式二
	p2 := Person {"zhang",20}
	// p2.Name = "tom"
	// p2.Age = 100
	fmt.Println(p2)

	// 方式三
	var p3 *Person = new(Person)

	// 因为 p3 是一个指针，因此标准的给字段赋值方式
	// 先用 * 取值符，将 p3 结构体变量找到，然后 . 里面的字段
	p3.Name = "p3"
	p3.Age = 10
	fmt.Println(*p3)

	// 方式四
	var person *Person = &Person{"p4大括号",20}

	// 因为 person 是一个指针，因此标准的访问字段的方法
	// (*person) 先通过 * 取值符找到我们的结构体，然后再 . Name 给他赋值
	// 但是 go 的设计者为了程序员使用方便，也可以 person.Name = "p4" ,原理和方式三 一样
	// person.Name = "p4" 
	// person.Age = 12
	fmt.Println(*person)

}