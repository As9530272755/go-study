package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 10)

	// 定义一个 map 类型数据 m1
	m1 := make(map[string]int)
	m1["1"] = 2

	// 定义 stu 数据类型为 Stu 结构体
	stu := Stu{
		Name: "tom",
		Age:  12,
	}

	// 定义 string 类型
	str := "这是一个 channel"

	// 定义一个 intr 的 int 类型
	intr := 10

	// 然后将所有定义的数据类型，全部写入到 allChan 管道中
	allChan <- m1
	allChan <- stu
	allChan <- str
	allChan <- intr

	// 取出数据
	GetChan1 := <-allChan
	GetChan2 := <-allChan
	GetChan3 := <-allChan
	GetChan4 := <-allChan

	fmt.Printf("chan1 = %v\nchan2 = %v\nchan3 = %v\nchan4 = %v\n", GetChan1, GetChan2, GetChan3, GetChan4)
}
