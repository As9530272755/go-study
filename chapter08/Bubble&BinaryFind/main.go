package main

import (
	"fmt"
)

func Find(arr *[6]int,leftIndex int , rightIndex int , findVal int)  {
	
	// 如果 leftIndex 下标大于 rightindex 就说明这个数组的下标已经遍历完了找不到了
	if leftIndex > rightIndex{
		fmt.Println("找不到")
		return
	}

	// 定义中间这个 mid 下标
	mid := leftIndex + (rightIndex - leftIndex) / 2

	// 如果中间下标 arr[mid] 大于我们要查找的值，说明我们查找得值一定在 mid 左边的数组
	if (*arr)[mid] > findVal {

		// 递归该函数，这里是将 mid-1 传递给了 Find(rightIndex) 这个形参，从而判断 if leftIndex > rightIndex
		Find(arr ,leftIndex,mid -1,findVal)
	
	} else if (*arr)[mid] < findVal {
		// 递归该函数，这里是将 mid+1 传递给了 Find(leftIndex) 这个形参，从而判断 if leftIndex > rightIndex 
		Find(arr,mid + 1 , rightIndex,findVal)
	
		// 这里判断的就是 mid = FindVal 的值，就代表找到了
	} else {
		fmt.Printf("已经找到 %v 下标为 %v",(*arr)[mid],mid)
	}

}

func main()  {

	// 定义一个乱序的数组
	arr := [...]int{11,3,42,31,2,192}

	// 定义一个自定义变量
	temp := 0

	// 编写一个冒泡排序，将 arr 数组排列整齐
	for j := 0 ; j < len(arr) -1 ; j++ {
		for i := 0 ; i < len(arr) -1 ; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
			}
		}
	}

	fmt.Println("冒泡排序之后",arr)
	
	// 调用 Find 自定义函数，然后 &arr 传入地址，0 为上面 leftIndex 值，len(arr) -1 为 rightIndex 值， 2 为 findVal 值
	Find(&arr,0,len(arr)-1,3)
}