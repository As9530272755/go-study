package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 写一个二分查找的 Find 函数
func Find(arr *[10]int,LeftIndex int,RighIndex int,FindVla int) {
	
	// 如果 leftindex 大于了 righindex 就提没有找到
	if LeftIndex > RighIndex {
		fmt.Println("没有找到")
		return
	}

	// 定义变量 mid 中间的值
	MidIndex := LeftIndex + (RighIndex - LeftIndex) / 2
	
	// 迭代循环该函数
	if (*arr)[MidIndex] < FindVla {
		Find(arr,MidIndex + 1,RighIndex,FindVla)
	} else if (*arr)[MidIndex] > FindVla {
		Find(arr,LeftIndex,MidIndex -1,FindVla)
	} else {
		fmt.Printf("找到 %v ,下标为 %v",(*arr)[MidIndex],MidIndex)
	}
}

func main()  {
	// 定义一个种子
	rand.Seed(time.Now().UnixNano())

	// 定义 arr 数组，并且将其赋值 10 个随机随机数
	var arr[10]int
	for i := 0 ; i < len(arr) ;i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("随机数赋值后",arr)

	// 冒泡排序
	temp := arr[0]
	for i := 0; i < len(arr)-1 ;i++ {
		for j := 0 ; j < len(arr)-1 ;j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println("排序后",arr)

	// 调用 Find 函数，并且传递 arr ,LeftIndex ,RighIndex ,FindVla 为&arr,0,len(arr)-1,90
	Find(&arr,0,len(arr)-1,90)
}