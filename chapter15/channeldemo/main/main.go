package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 2)

	intChan <- 110
	num := 20

	intChan <- num

	fmt.Printf("向管道写入数据后 channel 长度 = %v\n", len(intChan))

	num1 := <-intChan
	num2 := <-intChan
	fmt.Printf("向管道读取数据后 num1 = %v\n向管道读取数据后 num2 = %v",
		num1, num2)
}
