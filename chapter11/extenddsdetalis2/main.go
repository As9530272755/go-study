package main

import "fmt"

type A struct {
	Name string
	Age int
}

func (a *A) hello() {
	fmt.Printf("A name = %v A age = %v",a.Name,a.Age)
}

type C struct {
	a A
}

func main()  {
	b := C{}
	b.a.Name = "c.a"
	b.a.Age = 12
	b.a.hello()
}