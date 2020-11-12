package main

import (
	"fmt"
    "unsafe"
)

func main(){
	// var a int = 8900
	// var b uint = 1
	// // var c byte = 255
	// var n1 = 100
	// #fmt.Printf("n1 的 类型 %T",n1)

	var n2 = 10 
	fmt.Printf("n2 的类型 %T 占用的字节大小为 %d",n2,unsafe.Sizeof(n2))

	var age byte = 90
	fmt.Println("age=",age)
}