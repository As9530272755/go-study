package main

import (
	"fmt"
)

func main()  {
// 先写内循环再写外循环

	// 统计3个班成绩情况，每个班有5名同学
	// 求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]。

	// 分析思路1：
	// 1.先统计一个班的成绩情况，每个班有 5 名学生，求出该班的平均分。【先易后难】
	// 2.学生数就是 5 个 【先死后活】
	// 3.申明一个 sum 变量统计班级的总分

	// 分析思路2：
	// 1.统计3个班成绩情况，每个班有5个学生，求出每个班的平均分
	// 2.j 表示第几个班
	// 3.在定义一个变量来存放总成绩

    // 分析思路3：
	// 1. 我们可以将代码做活
	// 2. 定义两个变量，表示班级和人数
	
	
	// 分析思路4：
	// 1.我们给及格人数定义一个变量
	// 走代码
    
    var classNum int = 3
    var stuNum   int = 5
	var totalsum float64 = 0.0
	var pass int = 0 
	for j := 1 ; j <= classNum ; j++ {	
		sum := 0.0					// sum 的初始值为 0.0
		for i := 1 ; i <= stuNum ; i++ {	// 第一个循环我先写一个班的 5 名学生的平均分
			var score float64		// 给 score 变量定义一个为 float64 的类型
			fmt.Printf("请输入第%d班，第%d个学生的成绩:\t",j, i) // 通过 printf 引用 i 的值
			fmt.Scanln(&score)		// 通过 Scanln 调用上面的 score 函数
			
			if score >= 60 {
				pass++
			}
			// 累计分数
			sum += score			// sum 依次累计循环并且相加 score 的值
		}
		fmt.Printf("%d班级平均分是%v\n",j,sum / float64(stuNum)) 
		totalsum += sum
	}
	// 将各个班的总成绩累计到 totalsum中
	fmt.Printf("班级总分为%v,每个班的平均成绩为%v\n", totalsum, totalsum / float64(stuNum) * float64(classNum) )
	fmt.Printf("及格人数为%v",pass)
}