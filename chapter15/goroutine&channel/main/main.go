package main

import (
	"fmt"
	"strconv"
	"time"
)

// 编写一个函数，每隔 1 秒输出 "hello world"
func test() {
	for i := 0; i < 5; i++ {
		fmt.Println("test() hello world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

// 在 main 函数中输出 hello Golang
func main() {

	// go test()，表示对 test() 函数开启了协程
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() hello Golang " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
