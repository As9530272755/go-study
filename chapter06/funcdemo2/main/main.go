package main

import (
	"fmt"
)

func getSumAndSub(n1 int,n2 int) (int,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum,sub
}

func main()  {
	res1,res2 := getSumAndSub(20,10)	
	fmt.Printf("sum=%v\nsub=%v",res1,res2)
}