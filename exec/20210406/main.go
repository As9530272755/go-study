package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
	Colour string
}

func main()  {
	var name string
	fmt.Println("请输入猫猫的名字：")
	fmt.Scanln(&name)

	cat1 := Cat {
		Name : "小白",
		Age : 3,
		Colour : "白色",
	}
	cat2 := Cat {
		Name : "小花",
		Age : 100,
		Colour : "花色",
	}
	if name == cat1.Name {
		fmt.Printf("小猫名字 = %v 小猫年龄 = %v 小猫花色 = %v",
		cat1.Name,cat1.Age,cat1.Colour)
	} else if name == cat2.Name {
		fmt.Printf("小猫名字 = %v 小猫年龄 = %v 小猫花色 = %v",
		cat2.Name,cat2.Age,cat2.Colour)
	} else {
		fmt.Printf("张老太没有 %v 这只猫猫",name)
	}
}