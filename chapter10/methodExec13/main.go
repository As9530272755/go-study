package main

import (
	"fmt"
)

type Visitor struct {
	Name string
	Age int
}

func (v *Visitor) Price() {
	if v.Age > 18 {
		fmt.Printf("%v 年龄:%v 收费:20\n",v.Name,v.Age)
	} else {
		fmt.Printf("%v 年龄:%v 免费\n",v.Name,v.Age)
	}
}

func main()  {
	var v Visitor
	for {
	fmt.Println("请输入姓名：")
	fmt.Scanln(&v.Name)
	if v.Name == "n" {
		fmt.Println("退出程序...")
		break
	}
	fmt.Println("请输入年龄:")
	fmt.Scanln(&v.Age)
	v.Price()
	}
}