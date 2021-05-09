package main

import (
	"fmt"
	"runtime"
)

func main() {
	Cpu := runtime.NumCPU()
	fmt.Println("cpu 数量", Cpu)

	runtime.GOMAXPROCS(Cpu - 1)
	fmt.Println("ok")
}
