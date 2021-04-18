package main

import (
	"fmt"
)

type Monkey struct {
	Name string
}

func (m *Monkey) climb() {
	fmt.Println(m.Name,"会爬树...")
}

type Flight interface {
	Flight()
}

type Swim interface {
	Swim()
}

type LittleMonkey struct {
	Monkey
}

func (l *LittleMonkey) Flight() {
	fmt.Println(l.Name,"会飞翔...")
}

func (l *LittleMonkey) Swim() {
	fmt.Println(l.Name,"会游泳...")
}

func main()  {
	monkey := LittleMonkey {
		Monkey {
			Name : "孙悟空",
		},
	}
	monkey.climb()
	monkey.Flight()
	monkey.Swim()
}