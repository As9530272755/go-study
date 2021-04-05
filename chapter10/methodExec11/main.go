package main

import (
	"fmt"
)

type Dog struct {
	Name string
	Age int
	weight string
}

func (dog *Dog) Say() string {
	infoStr := fmt.Sprintf("Dog 信息：name = [%v] age = [%v] weight = [%v]",
	dog.Name,dog.Age,dog.weight)
	return infoStr
}

func main()  {
	var d Dog
	d.Name = "花花"
	d.Age = 3
	d.weight = "10斤"
	info := d.Say()
	fmt.Println(info)
}