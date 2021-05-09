package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	lock  sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= n
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 3)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("k = %d v = %d\n", i, v)
	}
	lock.Unlock()
}
