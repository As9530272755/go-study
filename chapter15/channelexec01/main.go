package main

import (
	"fmt"
)

var (
	ma = make(map[int]int)
)

func writeMap(n int, mapChan chan map[int]int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		ma[i] = res
	}

	for i := 1; i <= n; i++ {
		mapChan <- ma
	}

	close(mapChan)
}

func readMap(n int, mapChan chan map[int]int, exitChan chan bool) {
	for i := 0; i <= n; i++ {
		x, ok := <-mapChan
		if !ok {
			break
		}
		fmt.Printf("ma[%d] = %d\n", i, x[i])
	}

	exitChan <- true
	close(exitChan)
}

func main() {
	n := 200
	mapChan := make(chan map[int]int)
	exitChan := make(chan bool)
	go readMap(n, mapChan, exitChan)
	go writeMap(n, mapChan)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
