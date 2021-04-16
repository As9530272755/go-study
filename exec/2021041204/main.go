package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//要求：随机生成 5 个数，并将其反转打印，复杂应用
	var intArr[5]int = [...]int{1,2,3,4,5}
	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i < len(intArr) ; i++ {
		intArr[i] = rand.Intn(100)
		fmt.Println(intArr)
		for j := 0 ; j < len(intArr) / 2; j++ {
			intArr[j],intArr[len(intArr)-1-j] = intArr[len(intArr)-j-1],intArr[j]
		}
	}
	
	fmt.Println(intArr)
}