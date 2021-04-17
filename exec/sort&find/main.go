package main

import (
	"fmt"
)

func BubbleSort(arr []int){
	for j := 0 ; j < len(arr) - 1 ;j++{
		for i := 0 ; i < len(arr) -1 ; i++ {
			if arr[i] < arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func BinarySort(slice []int,left ,right ,find int)  {
	if left > right {
		fmt.Println("找不到",find)
		return
	}

	mid := left + (right - left) / 2
	
	if find < slice[mid] {
		BinarySort(slice, left, mid -1 ,find)
	} else if find > slice[mid] {
		BinarySort(slice, mid + 1 ,right,find)
	} else {
		fmt.Printf("找到了 %v 下标为 %v",find,mid)
		return
	}
	
}

func main() {
	arr := []int{24,69,80,57,13}
	BubbleSort(arr)

	hero := []string{"白眉鹰王","金毛狮王","紫衫龙王","青翼福王"}
	
	var find string
	fmt.Println("请输入你想要查找的 hero：")
	fmt.Scanln(&find)

	for i := 0 ; i < len(hero) ; i++ {
		if find == hero[i] { 
			fmt.Printf("已经找到 hero %v",find)
			break
		} else if i == len(hero) - 1{
			fmt.Println("没有找到",find)
			break
		}
	}
	slice := []int{1,8,10,89,1000,1234}
	BinarySort(slice,0,len(slice)-1,100000)
} 