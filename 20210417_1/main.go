package main

import (
	"fmt"
)

func main()  {
	arr := [5]int{1,22,33,44,55}
	var arr1[6]int
	arr1[len(arr)] = 21
	for i,_ := range arr {
		arr1[i] = arr[i]
	}
	
	for i := 0 ; i < len(arr1) -1; i++ {
		for j := 0 ; j < len(arr1) -1;j++ {
			if arr1[j] > arr1[j+1] {
				arr1[j],arr1[j+1] = arr1[j+1],arr1[j]
			}
		}
	}
	fmt.Println(arr1)
}