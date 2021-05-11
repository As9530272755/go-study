package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	map1 = make(map[int]int, 10)
	lock sync.Mutex
)

func test(n int) {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	// 这里是对 map1[n] 里面写数据，也就是将 res 阶乘的结果写入到 map1 中
	lock.Lock()
	map1[n] = res

	// 这里是对 map1 的对进行一个操作
	for i, v := range map1 {
		fmt.Printf("k = %v v = %v\n", i, v)
	}

	// 然后读完之后对 map1 的整个读写进行解锁机制
	lock.Unlock()
}

func main() {
	for i := 0; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

}
