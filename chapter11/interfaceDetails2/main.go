package main

import "fmt"

// 定义 A 接口，方法是 Say()
type A interface {
	Say()
}

// 定义 B 接口，方法是 Hello()
type B interface {
	Hello()
}

// 定义 Monster 结构体，并且想让这个 Monster 结构体既实现 A 接口中的 Say 方法，还有 B 接口中的 Hello 方法
type Monster struct {

}

// 下面定义了 Say 和 Hello 方法，此时此刻这个 Monster 结构体就是先了 A 和 B 接口

func (m Monster) Say() {
	fmt.Println("Say...")
}

func (m Monster) Hello() {
	fmt.Println("hello...")
}

func main()  {
	// 定义 Ms 变量类型为 Monster 结构体
	var Ms Monster 
	
	// 在定义 mon 变量类型为 A B 接口，将 Ms 结构体传递过去
	var mon A = Ms
	var mon1 B = Ms
	mon.Say()
	mon1.Hello()
} 