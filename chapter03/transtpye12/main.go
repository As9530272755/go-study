package main

import "fmt"

func main()  {
	
	var i int = 100
	var n1 float32 = float32(i)
	fmt.Printf("i=%v n1=%v\n",i,n1)

	fmt.Printf("i type is %T\n",i)

	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2)
}