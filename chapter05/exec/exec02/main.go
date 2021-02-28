package main

import (
	"fmt"
)

func main()  {
	var a float64 = 11.1
	var b float64 = 19.1
	if a > 10 && b < 20 {
		fmt.Println("a+b = ",a+b)
	}
}