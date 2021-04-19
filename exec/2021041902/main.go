package main

import (
	"fmt"
)

func main()  {
	arr := [8]int{1,2,3,100,5,6,7,8}
	
	sum := 0
	for i := 0 ; i < len(arr) ; i++ {
		sum += arr[i]
	}
	
	ave := sum / len(arr)
	max := 0
	mini := 0
	for i := 0 ; i < len(arr) ; i++ {
		if arr[i] > ave {
			max++
		} else if arr[i] < ave {
			mini++
		}
	}
	fmt.Printf("平均数：%v 大于平均数的个数：%v 小于平均数的个数：%v",
	ave,max,mini)
}