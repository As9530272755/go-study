package main

import (
	"fmt"
)

func main()  {

	var stu[3][5]int
	tot := 0
	for i := 0 ; i < len(stu) ; i++ {
		sum := 0
		for j := 0 ; j < len(stu[i]) ; j++ {
			fmt.Printf("请输入 %v 班第 %v 名学生成绩:\n",i+1,j+1)
			fmt.Scanln(&stu[i][j])
			sum += stu[i][j]
		}
		tot += sum
		fmt.Printf("%v 班平均分：%v\n",i + 1,sum / len(stu[i]))
	}
	fmt.Println("所有班级平均分：",tot / 15)
}