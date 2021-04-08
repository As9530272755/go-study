package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
}

type Stu Student

func main()  {
	var s1 Student
	var s2 Stu

	s1.Name = "zzz"
	s1.Age = 12

	s2 = Stu(s1)
	fmt.Printf("Name = %v Age = %v",s2.Name,s2.Age)
}