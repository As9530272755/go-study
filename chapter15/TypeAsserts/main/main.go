package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 12)
	cat1 := Cat{
		Name: "tom",
		Age:  12,
	}

	// 这里将 10 和 cat1 结构体传入到 allChan 管道中
	allChan <- 10
	allChan <- cat1

	// 我们只需要 allChan 管道中的第二个数据，所以将第一个数据 <-allChan 推走
	<-allChan

	// 然后再将第二个数据推给 cat11
	cat11 := <-allChan

	// 这里通过类型断言将 cat11 转换为 Cat 结构体
	newCat := cat11.(Cat)

	fmt.Println(newCat.Name)
}
