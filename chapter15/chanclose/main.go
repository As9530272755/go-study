package main

import "fmt"

func main() {
	// 定义了 intChan 的 channel 容量为 3
	intChan := make(chan int, 3)

	// 只往该 intChan 里面放入两个数据
	intChan <- 100
	intChan <- 200

	// 这里直接 close intChan ，将该 intChan 关闭
	close(intChan)

	// 当管道关闭后读取数据还是可以的
	fmt.Println("ok,test")
	num1 := <-intChan
	num2 := <-intChan
	fmt.Printf("num1 = %v\nnum2 = %v", num1, num2)
}
