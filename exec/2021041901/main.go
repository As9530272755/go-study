package main

import (
	"fmt"
)

func Arr(arr *[6]int)  {
	maxv := arr[0]
	maxi := 0
	miniv := arr[0]
	minii := 0
	for i := 0 ; i < len(arr) ;i++ {
		if maxv < arr[i] {
			maxv = arr[i]
			maxi = i
		} else if miniv > arr[i]{
			miniv = arr[i]
			minii = i
		}
	}
	fmt.Printf("最大值为 %v 最大下标 %v\n最小值为 %v 最小值下标 %v",
	maxv,maxi,miniv,minii)
}

func main()  {
	arr := [...]int{1,2,3,4,5,6}
	Arr(&arr)
}