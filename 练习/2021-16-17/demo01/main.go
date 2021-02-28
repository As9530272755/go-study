package main

import (
	"fmt"
)

func main()  {

for j := 1 ; j <= 3 ; j++{
		var stuNum int = 5
			sum := 0.0
		for i := 1 ; i <=stuNum ; i++ {
			var score float64
			fmt.Printf("请输入第%d 学生的成绩:",i)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("平均分为：%v",sum / float64(stuNum))
	}
}