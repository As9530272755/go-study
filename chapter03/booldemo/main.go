package main

import (
	"fmt"
	"unsafe"
)

func main ()  {
	var b = false
	fmt.Println("b=",unsafe.Sizeof(b))
}