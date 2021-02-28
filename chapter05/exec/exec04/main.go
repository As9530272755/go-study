package main

import (
	"fmt"
)

func main()  {
	var year int32 = 2020
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0{
		fmt.Printf("%v 是闰年",year)
	} else {
		fmt.Printf("%v 不是闰年",year)
	}
}