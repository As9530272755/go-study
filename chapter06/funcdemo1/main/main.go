package main

import (
	"fmt"
	util "go_code/chapter06/funcdemo1/utils"
)


func main()  {
	fmt.Println("utils.go中的Num1=",util.Num1)
	var n1 float64 = 10
	var n2 float64 = 2
	var operator byte = '/'
	result := util.Cal(n1,n2,operator)
	fmt.Println("result~=",result)

	n1 = 5
	n2 = 3
	operator = '*'
	result = util.Cal(n1,n2,operator)
	fmt.Println("result~=",result)
}