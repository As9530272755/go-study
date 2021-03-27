package main

import (
	"fmt"
)

/*跳水比赛，8个评委打分。
运动员的成绩是8个成绩取掉一个最高分，去掉一个最低分，剩下的6个分数的平均分就是最后得分。

使用一维数组实现如下功能:
(1）请把打最高分的评委和最低分的评委找出来。
(2)找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
*/	

func main()  {
	// 定义数组
	var arr[8]float64

	// 通过遍历该数组，在终端输入每个裁判的分数
	for i := 0 ; i < len(arr) ;i++ {
		fmt.Printf("请输入 %d 评委的分数\n",i+1)
		fmt.Scanln(&arr[i])
	}

	// 定义max 变量取出最高分
	max := arr[0]
	MaxIndex := 0

	// 定义 low 变量取出最低分
	low := arr[0]
	LowIndex := 0

	// 定义 sum 变量取出他们的所有分数
	sum := 0.0

	// 通过遍历 arr 数组将最高分和最低分的裁判找出来,(1）请把打最高分的评委和最低分的评委找出来
	for i := 0 ; i < len(arr); i++ {
		sum += arr[i]
		if max < arr[i] {
			max = arr[i]
			MaxIndex = i
		}
		if low > arr[i] {
			low = arr[i]
			LowIndex = i
		}
	}


	// 定义 ave 变量获取平均分
	ave := (sum - max - low) / 6

	fmt.Printf("评分最高分裁判 %v 是 %v 号裁判\n",max,MaxIndex+1)
	fmt.Printf("评分最低分裁判 %v 是 %v 号裁判\n",low,LowIndex+1)
	fmt.Printf("减掉最高分和最低分，最后得分 %.2f\n",ave)

} 