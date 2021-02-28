package main

import (
	"fmt"
)

var name  = "tom"

func test01()  {
	fmt.Println(name)
}

func test02()  {
	name = "jack"
	fmt.Println(name)
}

func main()  {
	fmt.Println(name)
	test01()
	test02()
	test01()
}