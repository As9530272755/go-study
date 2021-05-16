package main

import (
	"fmt"
	"sync"
)

var (
	res  int
	lock sync.Mutex
)

func writeNum(n int, numChan chan int) {

	for i := 0; i <= n; i++ {
		numChan <- i
	}
	close(numChan)
}

func readNum(n int, numChan chan int, exitChan chan bool) {
	fmt.Println("go")
	for {
		lock.Lock()
		v, ok := <-numChan
		if !ok {
			break
		}
		res = res + v
		lock.Unlock()
	}
	exitChan <- true

}

func main() {
	n := 200
	numChan := make(chan int)
	exitChan := make(chan bool, 1)
	go writeNum(n, numChan)
	for i := 0; i < 8; i++ {
		go readNum(n, numChan, exitChan)
	}

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

	fmt.Println("res = ", res)
	close(exitChan)
}
