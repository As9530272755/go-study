package main

import (
	"fmt"
)

type MethodUtils  struct {
}


func (ma *MethodUtils) Get() {
	arr := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	for i := 0 ; i < len(arr) ;i++ {
		for j := 0 ; j < len(arr[i]); j++ {
			temp := arr[0][0]
			if arr[0][1] == 2 {
				temp = arr[0][1]
				arr[0][1] = arr[1][0]
				arr[1][0] = temp
			}
			if arr[0][2] == 3 {
				temp = arr[0][2]
				arr[0][2] = arr[2][0]
				arr[2][0] = temp
			}
			if arr[1][2] == 6 {
				temp = arr[1][2]
				arr[1][2] = arr[2][1]
				arr[2][1] = temp
			}
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}

func main()  {
	var mu MethodUtils
	mu.Get()
}