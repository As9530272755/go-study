package main

import (
	"fmt"
)

type English interface {
	English()
}

type Athletes struct {
	Project string
}

type Basketball struct {
	Athletes
}

type Football struct {
	Athletes
}

func (f *Football) English() {
	fmt.Println(f.Project,"会英语...")
}

type Students struct {
	Stu string
}

type College struct {
	Students
}

func (c * College) English() {
	fmt.Println(c.Stu,"会英语...")
}

type Middle struct {
	Students
}

func main()  {
	football := Football{
		Athletes {
			Project : "足球运动员",
		},
	}
	football.English()

	college := College {
		Students{
			Stu : "大学生",
		},
	}
	college.English()
}