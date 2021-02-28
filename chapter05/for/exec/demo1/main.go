package main

import (
	"fmt"
)

func main()  {
	var sum = 0
	for i := 1 ; i <= 100 ; i++ {
		if i % 9 == 0 {
		fmt.Printf("%v ",i)
			sum += i
		}
	}
	fmt.Printf("sum=%v\n",sum)
}