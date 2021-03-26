package main

import (
	"fmt"
)

func main()  {
	var arr[5]int = [...]int{1,3,5,7,9}

	
	temp := arr[0]
	for i := 0 ;i < len(arr) -1; i++ {
		for j := 0 ;j < len(arr)-1 ; j++ {
			if arr[j] < arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}