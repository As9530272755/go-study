package model

import (
	"fmt"
)

type student struct {
	name string
	age int
}

func NewStudent() *student {
	return &student {

	}
}

func (stu student) Say() {
	stu.name = "zgy"
	stu.age = 12
	fmt.Println(stu.name,stu.age)
}