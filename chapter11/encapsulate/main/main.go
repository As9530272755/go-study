package main

import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main()  {
	p := model.NewPerson("test")
	p.SetSal(10000)
	p.SetAge(23)
	fmt.Println(*p)
	fmt.Println(p.Name,"年龄 = ",p.GetAge(),"工资 = ",p.GetSal())
}