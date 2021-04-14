package main

import (
	"fmt"
)

type A interface {
	Say()
}

type B interface {
	Hello()
}

type C interface {
	A
	B
	Cxi()
}

type ABC struct {

}

func (abc ABC) Say() {
	fmt.Println("say...")
}

func (abc ABC) Hello() {
	fmt.Println("hello...")
}

func (abc ABC) Cxi() {
	fmt.Println("Cxi...")
}

func main()  {
	var abc ABC
	var aabbcc C = abc

	aabbcc.Say()
	aabbcc.Hello()
	aabbcc.Cxi()
}