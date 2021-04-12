package main

import (
	"fmt"
)

type A struct {
	Name string
	Age int
}

func (a *A) SayOk() {
	fmt.Println("A SayOK name = ",a.Name)
}

func (a *A) age() {
	fmt.Println("A age = ",a.Age)
}

type B struct {
	Name string
	score int
}

func (b *B) SayOk() {
	fmt.Println("B SayOK name = ",b.Name)
}

func (b *B) Score() {
	fmt.Println("B Score = ",b.score)
}

type C struct {
	A
	B
}

func main()  {
	c := C{}
	c.A.Name = "A"
	c.B.Name = "B"
	c.Age = 12
	c.score = 99
	c.A.SayOk()
	c.B.SayOk()
	c.age()
	c.Score()
}