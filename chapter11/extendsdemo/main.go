package main

import (
	"fmt"
)

// 编写学生考试系统

// 小学生
type Pupil struct {
	Name string
	Age int
	Score int
}

// 编写 showinfo 方法,显示他的成绩
func (p *Pupil) ShowInfo() {
	fmt.Printf("学生名字：%v 年龄：%v 成绩：%v\n",p.Name,p.Age,p.Score)
}

// 编写 SetScore 设置成绩方法
func (p *Pupil) SetScore(score int) {
	// 将 main 函数中传递过来的 score 赋值给 p.Score
	p.Score = score
}

// 编写 testing 当前状态
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

// 大学生考试，除此之外还有研究生考试
type Graduate struct {
	Name string
	Age int
	Score int
}

func (p *Graduate) ShowInfo() {
	fmt.Printf("学生名字：%v 年龄：%v 成绩：%v\n",p.Name,p.Age,p.Score)
}

func (p *Graduate) SetScore(score int) {
	p.Score = score
}

func (p *Graduate) testing() {
	fmt.Println("小学生正在考试中...")
}

func main()  {

	// 定义 p 变量，类型为 Pupil 结构体，并给 Name、Age 字段赋值
	p := Pupil {
		Name : "tom",
		Age : 10,
	}

	// 调用 testing 方法
	p.testing()

	// 调用设置成绩方法，传递 90
	p.SetScore(90)

	// 调用 showinfo 方法
	p.ShowInfo()

	g := Graduate {
		Name : "大学生",
		Age : 10,
	}
	g.testing()
	g.SetScore(90)
	g.ShowInfo()
}