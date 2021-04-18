package main

import (
	"fmt"
)

func main()  {
	arr := [...]string{"AA","BB","CC","DD","AA","AA","EE","GG","MM","PP"}
	find := "AA"
	cot := 0
	for i := 0 ; i < len(arr) ; i ++ {
		if find == arr[i] {
			cot++
			fmt.Println(find,"下标为",i)
		}
	}
	fmt.Printf("有 %v 个 %v",cot,find)
}