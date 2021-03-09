package main

import (
	"fmt"
)

func main()  {
	var n1 float64 = 3
	var n2 float64 = 5
	var n3 float64 = 1
	var n4 float64 = 3.4
	var n5 float64 = 2
	var n6 float64 = 50
	total := n1 + n2 + n3 + n4 + n5 + n6
	average := fmt.Sprintf("%.2f",total / 6)
	fmt.Printf("总重量 = %v\n平均重量 = %v",total,average)

}