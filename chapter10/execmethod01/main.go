package main

import (
	"fmt"
)

// 给 person 结构体添加 jisuan 方法，可以计算从 1+..+1000 的结果。
type Person struct {
	Name string
}

func (person Person) speak()  {
	fmt.Println(person.Name,"是一个 goodMan")
}

func (person Person) jisuan()  {
	res := 0
	for i := 0 ; i <= 1000 ;i++ {
		res +=i
	}
	fmt.Println(person.Name,"计算的结果 = ",res)
}

func main()  {
	var p Person
	p.Name = "掌柜元"
	p.speak()
	p.jisuan()
}