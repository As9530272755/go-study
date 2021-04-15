package main

import (
	"fmt"
	"sort"
	//"math/rand"
)

type Students struct {
	Name string
	Age int
	Score float64
}

type Student []Students

func (stu Student) Len() int {
	return len(stu)
}

func (stu Student) Less(i,j int) (bool){
	return stu[i].Score > stu[j].Score
}

func (stu Student) Swap(i,j int) {
	stu[i],stu[j] = stu[j],stu[i]
}

func main()  {
	var student Student
		stu := Students {
			Name : "小明",
			Age : 18,
			Score : 77.1,
		}
		stu1 := Students {
			Name : "龙王",
			Age : 19,
			Score : 90.0,
		}
		stu2 := Students {
			Name : "鸣人",
			Age : 18,
			Score : 56.2,
		}
		stu3 := Students {
			Name : "路飞",
			Age : 15,
			Score : 32.0,
		}
		stu4 := Students {
			Name : "佐助",
			Age : 18,
			Score : 99.0,
		}
		student = append(student,stu,stu1,stu2,stu3,stu4)
	
	fmt.Println("排序前：")
	for _,v := range student {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println("排序后：")
	sort.Sort(student)
	for _,v := range student {
		fmt.Println("成绩排序后",v)
	}
}