package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	var arr[10]int
	rand.Seed(time.Now().UnixNano())

	for j := 0 ; j < len(arr) ;j++ {
		arr[j] = rand.Intn(100)
	}
	
	maxv := arr[0]
	maxi := 0
	miniv := arr[0]
	minii := 0
	
	for i := 0 ; i < len(arr); i++ {
		if maxv < arr[i] {
			maxv = arr[i]
			maxi = i
			
		}
		if miniv > arr[i] {
			miniv = arr[i]
			minii = i
		}
	}
	fmt.Printf("倒序前：%v\n最大值：%v 最大下标%v\n最小值：%v 最小下标：%v\n",arr,maxv,maxi,miniv,minii)

	find := 55
	for i := 0 ; i < len(arr); i++ {
		if find == arr[i] {
			fmt.Printf("找到 %v 下标为 %v\n",find,i)
			break
		} else if i == len(arr) -1{
			fmt.Println("没有找到：\n",find)
			break
		}
	}

	tot := 0	
	for i := 0 ; i < len(arr) -1; i++ {
		tot += arr[i+1]
		for j := 0 ; j < len(arr) -1 ; j++ {
			if arr[j] > arr[j+1] {
			arr[j],arr[j+1] = arr[j+1],arr[j]
			}
			
		}
	}
	fmt.Printf("倒序后：%v\n平均值：%v",arr,tot / len(arr))
}