package main

import (
	"fmt"
)

func main()  {

	// 1.定义一个二维数组
	//   [3] 代表 3 个班
	//   [5] 代表每个班级有 5 名学生
	var scores [3][5]float64

	// 2.循环输入成绩
	//   先遍历 scores 二维数组的长度
	for i := 0 ; i < len(scores) ;i++ {

		// 对第二层数组进行遍历，给每个学生输入成绩
		for j := 0 ; j < len(scores[i]) ; j++ {
			fmt.Printf("请输入第 %d 班 %d 学生的成绩\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	// 定义一个变量，用于累计所有班级的成绩
	total := 0.0
	// 遍历输入成绩后的这个二维数组统计平均分
	
	for i := 0 ; i < len(scores) ; i++ {
		
		// 定义 sum 变量用于累计各个班级的总分
		sum := 0.0

		for j := 0 ; j < len(scores[i]) ; j++ {
			
			// sum 就将 i 这个班级的总分统计出来了
			sum += scores[i][j]
		}
		fmt.Printf("第 %d 班级的总分 %v , 平均分为 %v \n",
		i+1,sum,sum / float64(len(scores[i])))
		
		// 将各个班级的成绩每次累计给 total
		total += sum
	}

	// 输出所有班级的平均分
	fmt.Printf("所有班级的总分是 = %v , 所有班级的平均分是 = %v  ",
	total,total / 15 )
}