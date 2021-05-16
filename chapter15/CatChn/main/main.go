package main

import (
	"fmt"
)

// 定义一个 Cat 结构体，有两个字段 Name 和 Age
type Cat struct {
	Name string
	Age  int
}

func main() {
	// 定义 CatChan 管道，类型为 cat ，容量为 10
	CatChan := make(chan *Cat, 10)

	// 顶一个 cat1 和 cat2 两个结构体变量，并且赋值
	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "jack", Age: 12}

	// 将 cat1 和 cat2 两个数据写入到 catChan 中
	CatChan <- &cat1
	CatChan <- &cat2

	// 定义 getcat1 和 getcat2 将 catchan 中的数据取出
	getCat1 := <-CatChan
	getCat2 := <-CatChan
	fmt.Printf("cat1 = %v\ncat2= %v", *getCat1, *getCat2)
}
