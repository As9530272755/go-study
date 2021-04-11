package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

func (s *Student) ShowInfo() {
	fmt.Printf("学生名字：%v 年龄：%v 成绩：%v\n",s.Name,s.Age,s.Score)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func (s *Student) GetSum(n1 int,n2 int) int {
	return n1 + n2
}

type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}

func main()  {
	pupil := &Pupil{}
	pupil.Student.Name = "黄二狗"
	pupil.Student.Age = 30
	pupil.Student.SetScore(99)
	pupil.testing()
	pupil.Student.ShowInfo()
	fmt.Println("小学生算术题结果：",pupil.Student.GetSum(5,5))
	
	fmt.Println()

	graduate := &Graduate{}
	graduate.Student.Name = "张桂元"
	graduate.Student.Age = 25
	graduate.Student.SetScore(99)
	graduate.testing()
	graduate.Student.ShowInfo()
	fmt.Println("小学生算术题结果：",graduate.Student.GetSum(5,5))
}
