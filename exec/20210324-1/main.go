package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 随机生成10个整数 (1_100的范围) 保存到数组，
	// 并倒序打印以及求平均值、求最大值和最大值的下标、并查找里面是否有 55
	var arr [10]int
	rand.Seed(time.Now().UnixNano())

	for i:= 0 ; i < len(arr) ; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Printf("定义的随机数组值 %v\n",arr)

	temp := 0
	for j := 0 ; j < len(arr) -1 ; j++ {
		for i := 0 ; i < len(arr)-1 ; i++ {
			if arr[i] < arr[i+1] {
				temp = arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}

	MaxVla := arr[0]
	MaxIndex := 0
	for i := 0 ; i < len(arr) ; i++ {
		if MaxVla < arr[i] {
			MaxVla = arr[i]
			MaxIndex = i
		}
	}

	sum := 0

	for i := 0 ; i < len(arr) ; i++ {
		sum += arr[i]
	}
	fmt.Println("arr 数组平均值 sum = ",sum)
	fmt.Printf("arr 数组平均值 sum = %.2f",float64(sum) / float64(len(arr)))

	fmt.Println("按照倒序排列之后 arr = ",arr)
	fmt.Printf("最大值 = %v 最大值小标为 = %v\n",MaxVla,MaxIndex)

	Find := 55
	for i := 0 ; i < len(arr) ; i++ {
		if Find == arr[i]{
			fmt.Printf("找到 %v ",Find)
			break
		} else if i == len(arr) -1 {
			fmt.Println("没有找到 ",Find)
			break
		}
	}
}